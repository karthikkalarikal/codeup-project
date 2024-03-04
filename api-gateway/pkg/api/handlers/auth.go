package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"

	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
)

type authHandlerImpl struct {
	client client.AuthClient
	utils  utils.Utils
}

func NewAuthHandler(client client.AuthClient, utils *utils.Utils) interfaces.AuthHandler {
	return &authHandlerImpl{
		client: client,
		utils:  *utils,
	}
}

func (u *authHandlerImpl) UserSignUp(e echo.Context) error {
	fmt.Println("inside user sign up handler 1")
	var user request.UserSignUpRequest

	if err := e.Bind(&user); err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	fmt.Println("user ", user)
	if user.FirstName == "" || user.Email == "" || user.Password == "" || user.ConfirmPassword == "" || user.Username == "" || user.LastName == "" {
		err := errors.New(customerrors.NoEmptyValueError)
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	if user.ConfirmPassword != user.Password {
		err := errors.New(customerrors.NoMatchingPasswordError)
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	fmt.Println("inside user sign up handler 2")
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(user); err != nil {
		err := errors.New(customerrors.ValidatorError + err.Error())
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	userCreated, err := u.client.UserSignUp(e, user)
	if err != nil {
		// err := errors.New(customerrors.ValidatorError + err.Error())
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	token, err := u.utils.GetTokenString(userCreated.ID)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	u.utils.WriteJSON(e, http.StatusCreated, []interface{}{userCreated, token})
	return nil

}

func (u *authHandlerImpl) UserSignIn(e echo.Context) error {
	fmt.Println("inside user sign up handler 1")
	var user request.UserSignInRequest

	if err := e.Bind(&user); err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	if user.Email == "" && user.Username == "" || user.Password == "" {
		err := errors.New(customerrors.NoEmptyValueError)
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	userSignedIn, err := u.client.UserSignIn(e, user)
	if err != nil {
		// err := errors.New(customerrors.ValidatorError + err.Error())
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	token, err := u.utils.GetTokenString(userSignedIn.ID)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	u.utils.WriteJSON(e, http.StatusCreated, []interface{}{userSignedIn, token})
	return nil
}
