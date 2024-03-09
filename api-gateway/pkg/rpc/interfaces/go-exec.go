package interfaces

import "github.com/labstack/echo/v4"

type GoCodeExecRPC interface {
	WriteGoCode(echo.Context, []byte) ([]byte, error)
}
