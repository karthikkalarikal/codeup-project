package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"

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
//	@Tags			admin
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
//	@Tags			admin
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
//	@Tags			admin
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
