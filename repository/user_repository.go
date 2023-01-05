package repository

import (
	"context"
	"database/sql"

	"github.com/satriowibowo1701/e-commorce-api/model"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.UserRegis) error
	Update(ctx context.Context, tx *sql.Tx, user model.UserUpdate) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (*model.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error)
}
