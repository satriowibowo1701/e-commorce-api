package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.UserRegis) error {
	SQL := "insert into usert(username,name,password,address,email,create_at,role) values ($1,$2,$3,$4,$5,$6,$7)"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Name, user.Password, user.Address, user.Email, time.Now(), "customer")
	return helper.IfError(err, "Error Creating User")
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.UserUpdate, id int) error {
	SQL := "update usert set name = $1,password = $2, address=$3,email=$4  where id = $5"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Password, user.Address, user.Email, id)

	return helper.IfError(err, "Error updating user")
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (*model.User, error) {
	SQL := "select id,username,name,password,role,address,email,create_at from usert where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.Role, &user.Address, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, errors.New("error Scan")
		}
		return &user, nil
	}
	return nil, errors.New("User not found")

}
func (repository *UserRepositoryImpl) FindByIdAdmin(ctx context.Context, tx *sql.Tx, userId int) (*model.UserAdminView, error) {
	SQL := "select id,username,name,role,address,email,create_at from usert where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	user := model.UserAdminView{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Role, &user.Address, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, errors.New("error Scan")
		}
		return &user, nil
	} else {
		return nil, errors.New("User not found")
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error) {
	SQL := "select id,username,name,password,role,address,email,create_at from usert where username = $1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.Role, &user.Address, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, errors.New("error Scan")
		}
		return &user, nil
	} else {
		return nil, errors.New("User not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.UserAll, error) {
	SQL := "select id,name from usert"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.UserAll
	for rows.Next() {
		user := model.UserAll{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		users = append(users, &user)
	}
	return users, nil
}
