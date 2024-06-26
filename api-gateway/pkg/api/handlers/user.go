package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userHandlerImp struct {
	user    interfaces.UserClient
	utils   utils.Utils
	goexec  interfaces.GoCodeExecClient
	payment interfaces.PaymentClient
}

func NewUserHandler(payment interfaces.PaymentClient, user interfaces.UserClient, utils *utils.Utils, goexec interfaces.GoCodeExecClient) handler.UserHandler {
	return &userHandlerImp{
		user:    user,
		utils:   *utils,
		goexec:  goexec,
		payment: payment,
	}
}

// Problem godoc
//
//	@Summary		View problems
//	@Description	View all problems code-up
//	@Tags			general
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]response.Problem
//	@Failure		400	{object}	[]response.Problem
//	@Failure		401	{object}	[]response.Problem
//	@Failure		404	{object}	[]response.Problem
//	@Failure		500	{object}	[]response.Problem
//	@Router			/user/view [get]
func (u *userHandlerImp) ViewAllProblems(e echo.Context) error {

	body, err := u.user.ViewAllProblems(struct{}{})
	if err != nil {
		return err
	}

	u.utils.WriteJSON(e, http.StatusCreated, body)
	return nil
}

// Problem godoc
//
//	@Summary		get one problems
//	@Description	get one problem to display
//	@Tags			Problem Execution Service
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Problem ID"
//	@Success		200	{object}	response.Problem
//	@Failure		400	{object}	response.Problem
//	@Failure		401	{object}	response.Problem
//	@Failure		404	{object}	response.Problem
//	@Failure		500	{object}	response.Problem
//	@Router			/user/problem/{id} [get]
func (u *userHandlerImp) GetOneProblemById(e echo.Context) error {
	fmt.Println("here in get one problem by id handler")
	problemId := e.Param("id")
	// problemId := "adfasdf"
	fmt.Println("problem id", problemId)
	objectId, err := primitive.ObjectIDFromHex(problemId)
	fmt.Println("object id", objectId)
	if err != nil {
		fmt.Println("error in id", problemId)
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err

	}

	var id request.GetOneProblemById
	id.ID = objectId
	body, err := u.user.GetProblemById(e, id)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	if body.Prime {
		if !e.Get("prime").(bool) {
			err := errors.New("this is a question for prime users only")
			u.utils.ErrorJson(e, err, http.StatusUnauthorized)
			return err
		}
	}

	u.utils.WriteJSON(e, http.StatusCreated, body)
	return nil
}

// Problem godoc
//
//	@Summary		Execute code
//	@Description	The code the user sent will be executed
//	@Tags			Problem Execution Service
//	@Accept			text/plain
//	@Produce		text/plain
//	@Security		BearerAuth
//	@Param			code	body		string					true	"Go code to execute"
//	@Success		200		{object}	string					"success"
//	@Failure		400		{object}	response.JsonResponse	"Bad Request"
//	@Failure		401		{object}	response.JsonResponse	"Unauthorized"
//	@Failure		403		{object}	response.JsonResponse	"Forbidden"
//	@Failure		500		{object}	response.JsonResponse	"Internal Server Error"
//
//	@Router			/user/go/exec [post]
func (u *userHandlerImp) WriteCode(e echo.Context) error {

	code := e.Request().Body
	if code == nil {
		err := errors.New("nil point error")
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return errors.New("nil point error")
	}

	body, err := io.ReadAll(code)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	fmt.Println("here 1")
	defer func() {
		fmt.Println("err", err)
	}()
	out, err := u.goexec.WriteGoCode(e, &body)

	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}
	// outString := string(out)
	// fmt.Println("out sting ", outString)

	e.Blob(http.StatusOK, "text/plain", *out)
	return nil
}

// Problem godoc
//
//	@Summary		Execute code
//	@Description	The code the user sent will be executed and the result will be given
//	@Tags			Problem Execution Service
//	@Accept			text/plain
//	@Produce		text/plain
//	@Security		BearerAuth
//	@Param			id		path		string					true	"Problem ID"
//	@Param			code	body		string					true	"Go code to execute"
//	@Success		200		{object}	string					"success"
//	@Failure		400		{object}	response.JsonResponse	"Bad Request"
//	@Failure		401		{object}	response.JsonResponse	"Unauthorized"
//	@Failure		403		{object}	response.JsonResponse	"Forbidden"
//	@Failure		500		{object}	response.JsonResponse	"Internal Server Error"
//	@Router			/user/go/{id} [post]
func (u *userHandlerImp) ExecuteGoCodyById(e echo.Context) error {
	code := e.Request().Body
	if code == nil {
		err := errors.New("nil point error")
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return errors.New("nil point error")
	}
	id := e.Param("id")
	if id == "" {
		err := errors.New("nil point error")
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return errors.New("nil point error")
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		err = errors.New("invalid id " + err.Error())
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}

	body, err := io.ReadAll(code)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	fmt.Println("here 1")
	defer func() {
		fmt.Println("err", err)
	}()
	out, err := u.user.ExecuteGoCodyById(e, request.SubmitCodeIdRequest{
		ID:   objectId,
		Code: body,
	})

	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}

	e.Blob(http.StatusOK, "text/plain", out)
	return nil
}

// Problem godoc
//
//	@Summary		Forget Password
//	@Description	to replace the password with new one
//	@Tags			User Panel
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			password	body		request.NewPassword		true	"new password"
//	@Success		200			{object}	response.JsonResponse	"success"
//	@Failure		400			{object}	response.JsonResponse	"Bad Request"
//	@Failure		401			{object}	response.JsonResponse	"Unauthorized"
//	@Failure		403			{object}	response.JsonResponse	"Forbidden"
//	@Failure		500			{object}	response.JsonResponse	"Internal Server Error"
//	@Router			/user/panal/password [put]
func (u *userHandlerImp) ForgetPassword(e echo.Context) error {
	fmt.Println("here in forget passoword")
	passwordStr := new(request.NewPassword)

	if err := e.Bind(passwordStr); err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	if passwordStr.ConfirmPassword != passwordStr.Password {
		err := customerrors.NoMatchingPasswordError.String()
		u.utils.ErrorJson(e, errors.New(err), http.StatusBadGateway)
		return errors.New(err)
	}

	if passwordStr.Password == "" {
		err := customerrors.NilPointError
		u.utils.ErrorJson(e, errors.New(err), http.StatusBadGateway)
		return errors.New(err)
	}

	body, err := u.user.ForgetPassword(e, request.ForgotPassword{
		Id:       e.Get("id").(int),
		Password: passwordStr.Password,
	})

	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}

	u.utils.WriteJSON(e, http.StatusOK, body)
	return nil

}

// Problem godoc
//
//	@Summary		Search Problem
//	@Description	Get Problem by difficulty or tags
//	@Tags			general
//	@Produce		json
//	@Param			searchBy	query		string					true	"tags, difficulty "
//	@Param			value		query		string					false	"Difficulty: easy, hard, medium. Tags:array, strigs,etc"
//	@Success		200			{object}	string					"success"
//	@Failure		400			{object}	response.JsonResponse	"Bad Request"
//	@Failure		401			{object}	response.JsonResponse	"Unauthorized"
//	@Failure		403			{object}	response.JsonResponse	"Forbidden"
//	@Failure		500			{object}	response.JsonResponse	"Internal Server Error"
//	@Router			/user/problem [get]
func (u *userHandlerImp) GetProblemBy(e echo.Context) error {
	fmt.Println("here get problem by")
	search := e.QueryParam("searchBy")
	value := e.QueryParam("value")
	if search == "" || value == "" {
		err := errors.New(customerrors.NilPointError)
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	fmt.Println("search ", search, " value ", value)
	body, err := u.user.GetProblemBy(e, request.SearchBy{
		Field:  search,
		Search: value,
	})
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}

	u.utils.WriteJSON(e, http.StatusOK, body)
	return nil

}

// render the payment frontend
func (u *userHandlerImp) Payment(e echo.Context) error {
	// Data to pass to the template
	data := map[string]interface{}{
		"Title": "Payment Page",
	}

	// Render the HTML template
	if err := e.Render(http.StatusOK, "terminal.page.gohtml", data); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type stripePayload struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// get payment intent, it gets called in the front end
func (u *userHandlerImp) GetPaymentIntent(e echo.Context) error {
	// payment-intent
	var payload stripePayload
	if err := e.Bind(&payload); err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	out, err := u.payment.Payment(e, request.Stripe{
		Amount:   amount,
		Currency: payload.Currency,
	})
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	e.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	e.Response().Write(out)
	// e.Blob(http.StatusOK, "Content-Type, application/json", out)

	return nil
}

func (u *userHandlerImp) PaymentSuccess(e echo.Context) error {
	fmt.Println("here in payment success")
	err := e.Request().ParseForm()
	if err != nil {
		log.Println(err)
		u.utils.ErrorJson(e, err)
		return err

	}

	cardHolder := e.Request().Form.Get("cardholder_name")
	email := e.Request().Form.Get("email")
	paymentIntent := e.Request().Form.Get("payment_intent")
	paymentMethod := e.Request().Form.Get("payment_method")
	paymentAmount := e.Request().Form.Get("payment_amount")
	paymentCurrency := e.Request().Form.Get("payment_currency")

	data := make(map[string]any)

	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency
	fmt.Println("data", data)

	err = u.user.MakePrime(e, email)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}

	if err := e.Render(http.StatusOK, "home.gohtml", data); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// unsubscribe

// Problem godoc
//
//	@Summary		Unsubscribe
//	@Description	Unsubscribe from prime membership
//	@Tags			User Panel
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	response.User	"success"
//	@Failure		400	{object}	response.User	"Bad Request"
//	@Failure		401	{object}	response.User	"Unauthorized"
//	@Failure		403	{object}	response.User	"Forbidden"
//	@Failure		500	{object}	response.User	"Internal Server Error"
//	@Router			/user/panal/unsubscribe [patch]
func (u *userHandlerImp) UnSubscrbe(e echo.Context) error {
	id := e.Get("id").(int)
	body, err := u.user.UnSubscrbe(e, id)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	u.utils.WriteJSON(e, http.StatusOK, body)
	return nil

}
