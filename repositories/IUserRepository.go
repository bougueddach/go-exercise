package repositories

import "go-exercise/entities"

type IUserRepository interface {
	FindById(id int) (entities.User, error)
	Save(user *entities.User) error
}
