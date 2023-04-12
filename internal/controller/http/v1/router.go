package v1

import (
	"clean-arch/internal/usecase"
	"clean-arch/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, u usecase.User) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v1")
	{
		newUserRoutes(h, u, l)
	}
}
