package repositories

import "go-exercise/src/entities"

var (
	u1    = entities.User{1, "omar", "bougueddach@gmail.com", "picture"}
	u2    = entities.User{2, "ayoub", "bougueddach@gmail.com", "picture"}
	u3    = entities.User{3, "allali", "bougueddach@gmail.com", "picture"}
	users = []entities.User{u1, u2, u3}
)

type UserRepository struct {
	UserRepository IUserRepository
}

func (repository *UserRepository) FindById(id int) (entities.User, error) {
	for i := range users {
		if users[i].Id == id {
			return users[i], nil
		}
	}
	return entities.User{}, nil // That's a bad hack, I should let the client know that I couldn't find it
}

func (repository *UserRepository) Save(user *entities.User) error {
	for i := range users {
		if users[i].Id == user.Id {
			users[i] = *user
			break
		}
	}
	//Save() should add it to the slice when not exist but we don't need it
	return nil
}
