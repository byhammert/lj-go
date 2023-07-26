package categories

import "github.com/google/uuid"

type InputBody struct {
	Name string `json:"name" validate:"required|string|min_len:3|max_len:30"`
}

type OutputCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type OutputCategories struct {
	Categories []OutputCategory `json:"categories"`
}
