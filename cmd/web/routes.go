package main

import (
	"github.com/gin-gonic/gin"
)

func (s *server) routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/api/v1")
	{
		v1.GET("/authors", s.ShowAuthorsHandler)
		v1.GET("/authors/:id", s.GetAuthorHandler)
		v1.POST("/authors", s.CreateAuthorsHandler)
	}

	return router
}
