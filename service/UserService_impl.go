package service

import (
	"context"
	"errors"
	"strings"

	"github.com/satriowibowo1701/e-commorce-api/helper"
	"github.com/satriowibowo1701/e-commorce-api/model"
)

func (service *InitService) CreateUser(ctx context.Context, request model.UserRegis) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}

	tx, _ := service.DB.Begin()
	_, err1 := service.UserRepository.FindByUsername(ctx, tx, request.Username)
	if err1 == nil {
		return helper.TxRollback(errors.New("error"), tx, "Username Already Taken")
	}
	password_hash, _ := helper.HashPassword(request.Password)
	request.Password = password_hash
	err3 := service.UserRepository.Create(ctx, tx, request)
	return helper.TxRollback(err3, tx, "Cannot Create Account")
}

func (service *InitService) UpdateUser(ctx context.Context, request model.UserUpdate, id int) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}
	tx, _ := service.DB.Begin()
	request.Password, _ = helper.HashPassword(request.Password)
	err2 := service.UserRepository.Update(ctx, tx, request, id)
	defer helper.TxRollback(err2, tx, "erroor")
	if err2 != nil {
		return err2
	}

	return nil

}

func (service *InitService) FindAllUser(ctx context.Context) ([]*model.UserAll, error) {
	tx, _ := service.DB.Begin()
	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, errors.New("No Data Found")
	}
	return users, nil
}

func (service *InitService) FindUserById(ctx context.Context, userid int) (*model.User, error) {
	if userid == -1 {
		return nil, errors.New("No Cookie Id Found")
	}
	tx, _ := service.DB.Begin()
	user, err := service.UserRepository.FindById(ctx, tx, userid)
	if err != nil {
		return nil, errors.New("No Data Found")
	}
	return user, nil
}

func (service *InitService) FindUserByUsername(ctx context.Context, username string) (*model.User, error) {
	tx, _ := service.DB.Begin()
	user, err := service.UserRepository.FindByUsername(ctx, tx, username)
	if err != nil {
		return nil, errors.New("No Data Found")
	}
	return user, nil
}

func (service *InitService) Login(ctx context.Context, req model.LoginRequest) (string, error) {

	tx, _ := service.DB.Begin()
	user, err := service.UserRepository.FindByUsername(ctx, tx, req.Username)
	defer helper.TxRollback(err, tx, "Error")
	if err != nil {
		return "", err
	}
	tokench := make(chan string)
	errch := make(chan error)
	defer close(errch)
	defer close(tokench)
	go func() {
		token, _ := helper.GenerateToken(user.ID, strings.ToLower(user.Role))
		tokench <- token
	}()
	go func() {
		err1 := helper.VerifyPassword(user.Password, req.Password)
		errch <- err1
	}()
	rtoken := <-tokench
	if err1 := <-errch; err1 != nil {
		return "", errors.New("Wrong password")
	}

	return rtoken, nil

}
