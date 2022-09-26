package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/api/v1")
	{
		v1.GET("/authors", s.ShowAuthorsHandler)
		v1.GET("/authors/:id", s.GetAuthorHandler)
		v1.POST("/authors", s.CreateAuthorHandler)
	}

	router.GET("/health", s.ShowHealthHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	return router
}
