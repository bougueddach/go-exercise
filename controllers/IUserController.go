package controllers

import "go-exercise/dtos"

type IUserController interface {
	GetUser(id int64) dtos.UserDTO
	UpdateUser(id int64, name string, email string, avatar string)
}
