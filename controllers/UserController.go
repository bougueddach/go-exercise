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

func (controller *UserController) UpdateUser(id int64, name string, email string, avatar string) {
	user, err := controller.FindById(id)
	if err != nil {
		panic(err)
	}
	user.Update(name, email, avatar)
	controller.Save(&user)
}
