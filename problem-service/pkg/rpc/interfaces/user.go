package interfaces

import "context"

type UserRPCProblem interface {
	ExecuteGoCode(context.Context, []byte) ([]byte, error)
}
