package entites

import "github.com/google/uuid"

type CategoryUsecaseContract interface {
	Create(name string) (Category, error)
	Delete(id uuid.UUID) error
	List() ([]Category, error)
	FindByID(id uuid.UUID) (Category, error)
	Update(id uuid.UUID, name string) (Category, error)
}
