package category

import (
	entities "github.com/byhammert/lj-go/entities/categories"
	"github.com/google/uuid"
)

func (usecase *CategoryUsecase) FindByID(id uuid.UUID) (entities.Category, error) {
	return usecase.Database.CategoryRepository.FindByID(id)
}
