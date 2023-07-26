package api

import (
	"fmt"

	"github.com/byhammert/lj-go/infra/config"
	"github.com/byhammert/lj-go/infra/database"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Engine *gin.Engine

	Database *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{
		Engine:   gin.Default(),
		Database: db,
	}
}

func (s *Service) Start() error {
	s.GetRoutes()

	return s.Engine.Run(fmt.Sprintf(":%d", config.Env.ApiPort))
}
