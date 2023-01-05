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
	SQL := "insert into user(username,name,password,adress,email,create_at,role) values (?,?,?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Name, user.Password, user.Address, user.Email, time.Now(), "customer")
	return helper.TxRollback(err, tx, "Error Creating User")
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.UserUpdate) error {
	SQL := "update category set name = ?,password = ?, address=?=?,email=?  where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Password, user.Address, user.Email, user.Id)

	return helper.TxRollback(err, tx, "Error updating user")
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (*model.User, error) {
	SQL := "select id,username,name,password,role,adress,email,create_at from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer rows.Close()
	defer tx.Commit()
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

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (*model.User, error) {
	SQL := "select id,username,name,password,role,adress,email,create_at from user where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		return nil, errors.New("Error Sql")
	}
	defer tx.Commit()
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

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error) {
	SQL := "select id,username,name,password,role,adress,email,create_at from user"
	rows, err := tx.QueryContext(ctx, SQL)
	defer tx.Commit()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.Role, &user.Address, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, errors.New("Cannot Scaning")
		}
		users = append(users, &user)
	}
	return users, nil
}
