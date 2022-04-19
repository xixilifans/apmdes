package bean

type User struct {
	id       int
	name     string
	password string
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}
