package repositories

import "go-exercise/entities"

type IUserRepository interface {
	FindById(id int64) (entities.User, error)
}
