package service

import (
	"context"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type UserService_impl interface {
	CreateUser(ctx context.Context, request model.UserRegis) error
	UpdateUser(ctx context.Context, request model.UserUpdate) error
	FindAllUser(ctx context.Context) ([]*model.User, error)
	FindUserById(ctx context.Context, userid int64) (*model.User, error)
	Login(ctx context.Context, req model.LoginRequest) (string, error)
	FindUserByUsername(ctx context.Context, username string) (*model.User, error)
}
