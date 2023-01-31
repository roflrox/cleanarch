package usecases

import "context"

type UseCase interface {
	SayHello(ctx context.Context) (message string, err error)
}
