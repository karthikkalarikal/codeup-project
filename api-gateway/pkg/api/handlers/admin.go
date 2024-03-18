package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type adminHandlerImpl struct {
	client client.AdminClient
	utils  utils.Utils
}

func NewAdminHandler(client client.AdminClient, utils *utils.Utils) interfaces.AdminHandler {
	return &adminHandlerImpl{
		client: client,
		utils:  *utils,
	}
}

// Problem godoc
//
//	@Summary		Create a Problem
//	@Description	Admin create a problem
//	@Tags			Problem Management
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			user	body		request.InsertProblem	true	"create problem"
//	@Success		201		{object}	response.JsonResponse	"Success: Problem created"
//	@Failure		400		{object}	response.JsonResponse	"Bad request"
//	@Failure		401		{object}	response.JsonResponse	"Unauthorized"
//	@Failure		500		{object}	response.JsonResponse	"Internal server error"
//	@Router			/admin/problem/ [post]
func (a *adminHandlerImpl) CreateProblem(e echo.Context) error {
	fmt.Println("here in creat a problem handler")
	var problem request.InsertProblem

	if err := e.Bind(&problem); err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	if problem.Description == "" && problem.Difficulty == "" || problem.Title == "" {
		err := errors.New(customerrors.NoEmptyValueError.String())
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	body, err := a.client.InsertProblem(e, problem)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	a.utils.WriteJSON(e, http.StatusCreated, body)
	return nil
}

// Problem godoc
//
//	@Summary		Insert first half of problem
//	@Description	Admin insert first half of problem
//	@Tags			Problem Management
//	@Accept			text/plain
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			id		path		string					true	"Problem ID"
//	@Param			code	body		string					true	"insert into problem"
//	@Success		201		{object}	response.JsonResponse	"Success: Problem Modified"
//	@Failure		400		{object}	response.JsonResponse	"Bad request"
//	@Failure		401		{object}	response.JsonResponse	"Unauthorized"
//	@Failure		500		{object}	response.JsonResponse	"Internal server error"
//	@Router			/admin/problem/first/{id} [put]
func (a *adminHandlerImpl) InsertFirstHalfProblem(e echo.Context) error {
	fmt.Println("insert first half of problem")

	problemId := e.Param("id")
	// problemId := "adfasdf"
	fmt.Println("problem id", problemId)
	objectId, err := primitive.ObjectIDFromHex(problemId)
	fmt.Println("object id", objectId)
	if err != nil {
		fmt.Println("error in id", problemId)
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err

	}
	var problem request.FirstHalfCode

	code := e.Request().Body
	// code := e.Request().Body
	if code == nil {
		err := errors.New("nil point error")
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return errors.New("nil point error")
	}
	body, err := io.ReadAll(code)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	problem.ID = objectId
	problem.FirstHalfCode = body

	out, err := a.client.InsertFirstHalfProblem(e, problem)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	a.utils.WriteJSON(e, http.StatusCreated, out)
	return nil
}

// Problem godoc
//
//	@Summary		Insert second half of problem
//	@Description	Admin insert second half of problem
//	@Tags			Problem Management
//	@Accept			text/plain
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			id		path		string					true	"Problem ID"
//	@Param			code	body		string					true	"Modified problem"
//	@Success		201		{object}	response.JsonResponse	"Success: Problem Modified"
//	@Failure		400		{object}	response.JsonResponse	"Bad request"
//	@Failure		401		{object}	response.JsonResponse	"Unauthorized"
//	@Failure		500		{object}	response.JsonResponse	"Internal server error"
//	@Router			/admin/problem/second/{id} [put]
func (a *adminHandlerImpl) InsertSecondHalfProblem(e echo.Context) error {
	fmt.Println("insert second half of problem")

	var problem request.SecondHalfCode

	problemId := e.Param("id")

	fmt.Println("problem id", problemId)

	objectId, err := primitive.ObjectIDFromHex(problemId)
	if err != nil {
		fmt.Println("error in id", problemId)
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	code := e.Request().Body
	// code := e.Request().Body
	if code == nil {
		err := errors.New("nil point error")
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return errors.New("nil point error")
	}
	body, err := io.ReadAll(code)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	problem.ID = objectId
	problem.SecondHalfCode = body

	out, err := a.client.InsertSecondHalfProblem(e, problem)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	a.utils.WriteJSON(e, http.StatusCreated, out)
	return nil
}

// Problem godoc
//
//	@Summary		Get Users
//	@Description	Admin Gets the list of all users
//	@Tags			User Management
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Success		201	{object}	[]response.User	"Success: get all users"
//	@Failure		400	{object}	[]response.User	"Bad request"
//	@Failure		401	{object}	[]response.User	"Unauthorized"
//	@Failure		500	{object}	[]response.User	"Internal server error"
//	@Router			/admin/user/ [get]
func (a *adminHandlerImpl) ViewUsers(e echo.Context) error {
	out, err := a.client.ViewUsers(e)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	a.utils.WriteJSON(e, http.StatusCreated, out)
	return nil
}

// User godoc
//
//	@Summary		Search Users
//	@Description	Admin Gets the list of all users using keyworkd
//	@Tags			User Management
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			keyword	path		string			false	"keyword"
//	@Success		201		{object}	[]response.User	"Success: get all users"
//	@Success		204		{object}	[]response.User	"no users"
//	@Failure		400		{object}	[]response.User	"Bad request"
//	@Failure		401		{object}	[]response.User	"Unauthorized"
//	@Failure		500		{object}	[]response.User	"Internal server error"
//	@Router			/admin/user/{keyword} [get]
func (a *adminHandlerImpl) SearchUser(e echo.Context) error {

	keyword := e.Param("keyword")
	fmt.Println("keyword ", keyword)
	if keyword == "{keyword}" {
		keyword = ""
	}
	out, err := a.client.SearchUser(e, request.Search{
		Keyword:  keyword,
		SearchBy: "email",
	})

	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	if out == nil {
		return a.utils.WriteJSON(e, http.StatusNoContent, out)
	}

	a.utils.WriteJSON(e, http.StatusOK, out)
	return nil
}

// User godoc
//
//	@Summary		Block user
//	@Description	Admin Can block/unblock user by passing id
//	@Tags			User Management
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			id	path		int						true	"user id"
//	@Success		201	{object}	response.BlockedStatus	"Success: get all users"
//	@Success		204	{object}	response.BlockedStatus	"no users"
//	@Failure		400	{object}	response.BlockedStatus	"Bad request"
//	@Failure		401	{object}	response.BlockedStatus	"Unauthorized"
//	@Failure		500	{object}	response.BlockedStatus	"Internal server error"
//	@Router			/admin/user/{id} [patch]
func (a *adminHandlerImpl) BlockUser(e echo.Context) error {
	idStr := e.Param("id")
	fmt.Println("id ", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	fmt.Println("id ", id)

	out, err := a.client.BlockUser(e, id)
	if err != nil {
		a.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	a.utils.WriteJSON(e, http.StatusAccepted, out)
	return nil
}
