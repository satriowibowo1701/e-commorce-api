package service

import (
	"context"
	"errors"

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
	if err1.Error() != "User not found" {
		return err1
	}
	password_hash, _ := helper.HashPassword(request.Password)
	request.Password = password_hash
	err2 := service.UserRepository.Create(ctx, tx, request)
	if err2 != nil {
		return err2
	}
	return nil
}

func (service *InitService) UpdateUser(ctx context.Context, request model.UserUpdate) error {

	err := service.Validate.Struct(request)
	if err != nil {
		return err
	}

	tx, _ := service.DB.Begin()

	_, err1 := service.ProdukRepostory.FindById(ctx, tx, int(request.Id))
	if err1 != nil {
		return err1
	}

	err2 := service.UserRepository.Update(ctx, tx, request)
	if err2 != nil {
		return err2
	}
	return nil

}

func (service *InitService) FindAllUser(ctx context.Context) ([]*model.User, error) {
	tx, _ := service.DB.Begin()
	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, errors.New("No Data Found")
	}
	return users, nil
}

func (service *InitService) FindUserById(ctx context.Context, userid int64) (*model.User, error) {
	tx, _ := service.DB.Begin()
	user, err := service.UserRepository.FindById(ctx, tx, int(userid))
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
	if err != nil {
		return "", err
	}
	err1 := helper.VerifyPassword(user.Password, req.Password)
	if err1 != nil {
		return "", errors.New("Password Salah")
	}

	token, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil

}
