package Object

type User struct {
	Id,Age int
	Name string
}

func NewUser(f func(u *User)) *User {
	u := new(User)
	f(u)
	return u
}

func WithUserId(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}
}