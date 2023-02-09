package service

import (
	"github.com/satriowibowo1701/simple-rest-api/helper"
	"github.com/satriowibowo1701/simple-rest-api/model"
)

type UserService struct {
}

func NewUserService() UserServiceImpl {
	return &UserService{}
}
func (user *UserService) GetUsers() ([]model.User, error) {
	var res []model.User
	err := helper.ApiCall("GET", "https://api.github.com/users", nil, &res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (user *UserService) GetDetailUser(nama string) (*model.UserDetail, error) {

	var res model.UserDetail
	err := helper.ApiCall("GET", "https://api.github.com/users/"+nama, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil

}
