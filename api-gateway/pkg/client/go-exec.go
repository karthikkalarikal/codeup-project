package client

import (
	"fmt"

	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/labstack/echo/v4"
)

type goCodeExecClientImpl struct {
	client client.GoCodeExecClient
}

func NewGoExecClient(client interfaces.GoCodeExecRPC) client.GoCodeExecClient {
	return &goCodeExecClientImpl{
		client: client,
	}
}

func (c *goCodeExecClientImpl) WriteGoCode(e echo.Context, data *[]byte) (*[]byte, error) {
	fmt.Println("here 2")
	body, err := c.client.WriteGoCode(e, data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
