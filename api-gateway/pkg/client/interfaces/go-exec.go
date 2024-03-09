package interfaces

import "github.com/labstack/echo/v4"

type GoCodeExecClient interface {
	WriteGoCode(echo.Context, *[]byte) (*[]byte, error)
}
