package controllers

import (
	"go-exercise/entities"
	"go-exercise/repositories"
)

type UserController struct {
	repositories.IUserRepository
}

func (controller *UserController) GetUser(id int64) (entities.User) {
	user, err := controller.FindById(id)
	if err != nil {
		panic(err)
	}
	return user
}