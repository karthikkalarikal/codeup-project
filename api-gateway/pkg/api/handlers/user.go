package handlers

import (
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
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

func (u *userHandlerImp) ViewAllProblems(e echo.Context) error {

	body, err := u.user.ViewAllProblems(struct{}{})
	if err != nil {
		return err
	}

	u.utils.WriteJSON(e, http.StatusCreated, body)
	return nil
}
