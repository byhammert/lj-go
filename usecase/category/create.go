package category

import entites "github.com/byhammert/lj-go/entities/categories"

func (usecase *CategoryUsecase) Create(name string) (entites.Category, error) {
	category := entites.NewCategory(name)

	err := usecase.Database.CategoryRepository.Create(category)

	return *category, err
}
