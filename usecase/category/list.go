package category

import (
	entities "github.com/byhammert/lj-go/entities/categories"
)

func (usecase *CategoryUsecase) List() ([]entities.Category, error) {
	return usecase.Database.CategoryRepository.List()
}
