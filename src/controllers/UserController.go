package controllers

import (
	"go-exercise/src/dtos"
	"go-exercise/src/repositories"
)

type UserController struct {
	repositories.IUserRepository
}

func (controller *UserController) GetUser(id int) dtos.UserDTO {
	user, err := controller.FindById(id)
	if err != nil {
		panic(err)
	}
	return dtos.UserDTO{user.Id, user.Name, user.Email, user.Avatar}
}

func (controller *UserController) UpdateUser(id int, name string, email string, avatar string) {
	user, err := controller.FindById(id)
	if err != nil {
		panic(err)
	}
	user.Update(name, email, avatar)
	controller.Save(&user)
}
