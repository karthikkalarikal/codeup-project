package handlers

import (
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/labstack/echo/v4"
)

type userHandlerImp struct {
	user  interfaces.UserClient
	utils utils.Utils
}

func NewUserHandler(user interfaces.UserClient, utils *utils.Utils) handler.UserHandler {
	return &userHandlerImp{
		user:  user,
		utils: *utils,
	}
}

// Problem godoc
//
//	@Summary		View problems
//	@Description	View all problems code-up
//	@Tags			user,admin
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

func (u *userHandlerImp) GetOneProblemById(e echo.Context) error {

	var id request.GetOneProblemById
	if err := e.Bind(&id); err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	body, err := u.user.GetProblemById(e, id)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}
	u.utils.WriteJSON(e, http.StatusCreated, body)
	return nil
}
