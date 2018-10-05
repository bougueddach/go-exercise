package controllers

import (
	"go-exercise/dtos"
	"go-exercise/repositories"
)

type UserController struct {
	repositories.IUserRepository
}

func (controller *UserController) GetUser(id int64) dtos.UserDTO {
	user, err := controller.FindById(id)
	if err != nil {
		panic(err)
	}
	return dtos.UserDTO{user.Id, user.Name, user.Email, user.Avatar}
}
