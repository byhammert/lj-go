package infra

import (
	"net/http"

	"github.com/byhammert/lj-go/api/controller"
	"github.com/gin-gonic/gin"
)

func Heart(c *gin.Context) {
	c.JSON(http.StatusOK, controller.NewResponseMessage("ok"))
}
