package controllers

import "go-exercise/dtos"

type IUserController interface {
	GetUser(id int) dtos.UserDTO
	UpdateUser(id int, name string, email string, avatar string)
}
