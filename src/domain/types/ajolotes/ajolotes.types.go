package types

import types "Gin/src/domain/types/task"

type Ajolotes struct {
	Nombre string       `json:"nombre"`
	Color  string       `json:"color"`
	Task   []types.Task `json:"task"`
}
