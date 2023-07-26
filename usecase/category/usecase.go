package category

import "github.com/byhammert/lj-go/infra/database"

type CategoryUsecase struct {
	Database *database.Database
}

func NewCategoryUsecase(db *database.Database) *CategoryUsecase {
	return &CategoryUsecase{
		Database: db,
	}
}
