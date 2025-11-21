package users

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Tasks    []Task
}
