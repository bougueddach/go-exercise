package controllers

import "go-exercise/dtos"

type IUserController interface {
	GetUser(id int64) dtos.UserDTO
}
