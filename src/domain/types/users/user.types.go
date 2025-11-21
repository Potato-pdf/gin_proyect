package types

import types "Gin/src/domain/types/task"

type User struct {
	Name     string
	Email    string
	Password string
	Tasks    []types.Task
}
