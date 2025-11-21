package types

import "Gin/src/domain/models/task"

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Tasks    []task.Task
}
