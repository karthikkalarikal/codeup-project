package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
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
		err := errors.New(customerrors.NoEmptyValueError)
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
