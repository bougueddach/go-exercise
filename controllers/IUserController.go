package controllers

import "go-exercise/entities"

type IUserController interface {
	GetUser(id int64) (entities.User)
}