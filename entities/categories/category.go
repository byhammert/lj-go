package entites

import (
	"github.com/byhammert/lj-go/entities/shared"
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name" bson:"_name"`
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   shared.GetUuid(),
		Name: name,
	}
}

type CategoryRepository interface {
	Create(category *Category) error
	List() ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
	Update(category *Category) error
	Delete(id uuid.UUID) error
}
