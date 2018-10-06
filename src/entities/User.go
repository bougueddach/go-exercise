package entities

type User struct {
	Id     int
	Name   string
	Email  string
	Avatar string
}

func (user *User) Update(name string, email string, avatar string) {
	user.Name = name
	user.Email = email
	user.Avatar = avatar
}
