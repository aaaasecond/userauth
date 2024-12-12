package biz

import (
	"context"
 //       "userauth/internal/data"
        "github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"
)

var ProviderSet = wire.NewSet(NewUserUseCase)

type UserRepo interface {
	AddUser(ctx context.Context, username, password string) error
	AuthenticateUser(ctx context.Context, username, password string) (bool, error)
	DeleteUser(ctx context.Context, username string) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) AddUser(ctx context.Context, username, password string) error {
	return uc.repo.AddUser(ctx, username, password)
}

func (uc *UserUseCase) AuthenticateUser(ctx context.Context, username, password string) (bool, error) {
	return uc.repo.AuthenticateUser(ctx, username, password)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, username string) error {
	return uc.repo.DeleteUser(ctx, username)
}
