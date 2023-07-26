package category

import (
	"errors"

	entities "github.com/byhammert/lj-go/entities/categories"
	"github.com/byhammert/lj-go/entities/shared"
	"github.com/google/uuid"
)

func (usecase *CategoryUsecase) Update(id uuid.UUID, name string) (entities.Category, error) {
	category, err := usecase.Database.CategoryRepository.FindByID(id)
	if err != nil {
		return category, err
	}

	if category.ID == shared.GetUuidEmpty() {
		return category, errors.New("Não foi possivel encontrar a categoria")
	}

	category.Name = name

	err = usecase.Database.CategoryRepository.Update(&category)

	return category, err
}
