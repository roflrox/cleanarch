package usecases

import (
	"cleanarch/domain"
	"context"
	"encoding/json"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func (u *UserUsecase) SayHello(ctx context.Context) (message string, err error) {
	user, _ := u.repo.GetByID(ctx, 1)
	userBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	message = string(userBytes)
	return
}

func NewUserUsecase(r domain.UserRepository) UseCase {
	return &UserUsecase{repo: r}
}
