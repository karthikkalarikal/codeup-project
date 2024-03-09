package handlers

import (
	"io"
	"net/http"

	handler "github.com/karthikkalarikal/api-gateway/pkg/api/handlers/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
	"github.com/labstack/echo/v4"
)

type userHandlerImp struct {
	user   interfaces.UserClient
	utils  utils.Utils
	goexec interfaces.GoCodeExecClient
}

func NewUserHandler(user interfaces.UserClient, utils *utils.Utils, goexec interfaces.GoCodeExecClient) handler.UserHandler {
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

// Problem godoc
//
//	@Summary		Execute code
//	@Description	The code the user sent will be executed by user
//	@Tags			user
//	@Accept			text/plain
//	@Produce		text/plain
//	@Success		200	{object}	[]response.Problem
//	@Failure		400	{object}	[]response.Problem
//	@Failure		401	{object}	[]response.Problem
//	@Failure		404	{object}	[]response.Problem
//	@Failure		500	{object}	[]response.Problem
//	@Router			/user/go/exec [post]
func (u *userHandlerImp) WriteCode(e echo.Context) error {
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadRequest)
		return err
	}

	out, err := u.goexec.WriteGoCode(e, body)

	if err != nil {
		u.utils.ErrorJson(e, err, http.StatusBadGateway)
		return err
	}
	// outString := string(out)
	// fmt.Println("out sting ", outString)

	e.Blob(http.StatusOK, "text/plain", out)
	return nil
}
