package api

import (
	"fmt"

	"github.com/byhammert/lj-go/api/controller/categories"
	"github.com/byhammert/lj-go/infra/config"
	"github.com/byhammert/lj-go/infra/database"
	"github.com/gin-gonic/gin"

	category_usecase "github.com/byhammert/lj-go/usecase/category"
)

type Service struct {
	Engine *gin.Engine

	Database           *database.Database
	CategoryController *categories.CategoryController
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) GetControllers() {
	categoryUsecase := category_usecase.NewCategoryUsecase(s.Database)
	s.CategoryController = categories.NewCategoryController(categoryUsecase)
}

func (s *Service) Start() error {
	s.GetRoutes()

	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
