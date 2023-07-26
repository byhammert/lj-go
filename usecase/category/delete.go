package category

import (
	"errors"

	"github.com/byhammert/lj-go/entities/shared"
	"github.com/google/uuid"
)

func (usecase *CategoryUsecase) Delete(id uuid.UUID) error {
	category, err := usecase.Database.CategoryRepository.FindByID(id)
	if err != nil {
		return err
	}

	if category.ID == shared.GetUuidEmpty() {
		return errors.New("Não foi possivel encontrar a categoria")
	}

	return usecase.Database.CategoryRepository.Delete(id)
}
