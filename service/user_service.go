package service

import "github.com/satriowibowo1701/simple-rest-api/model"

type UserServiceImpl interface {
	GetUsers() ([]model.User, error)
	GetDetailUser(nama string) (*model.UserDetail, error)
}
