package api

import (
	infra_controller "github.com/byhammert/lj-go/api/controller/infra"
)

func (s *Service) GetRoutes() {
	s.Engine.GET("/heart", infra_controller.Heart)

	groupCategory := s.Engine.Group("categories")
	groupCategory.GET("/", s.CategoryController.List)
	groupCategory.POST("/", s.CategoryController.Create)
	groupCategory.PUT("/:id", s.CategoryController.Update)
	groupCategory.DELETE("/:id", s.CategoryController.Delete)
	groupCategory.GET("/:id", s.CategoryController.Details)
}
