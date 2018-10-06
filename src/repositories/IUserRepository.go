package repositories

import "go-exercise/src/entities"

type IUserRepository interface {
	FindById(id int) (entities.User, error)
	Save(user *entities.User) error
}
