package main

import (
	"github.com/alexhiggins/go-api/internal/author"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

type server struct {
	queries  queries
	commands commands
	logger   *zap.Logger
	router   *gin.Engine
}

type commands struct {
	createAuthor author.CreateAuthorCommand
}

type queries struct {
	fetchAuthor     author.FetchAuthorQuery
	fetchAllAuthors author.FetchAllAuthorsQuery
}

func (s *server) run(port string) error {
	r := s.routes()

	if err := r.Run(port); err != nil {
		s.logger.Error("there was an error calling Run on router: %v", zap.Error(err))
		return err
	}

	return nil
}

func (s *server) okResponse(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func (s *server) notFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func (s *server) createdResponse(c *gin.Context, body interface{}) {
	c.JSON(http.StatusCreated, body)
}

func (s *server) failedValidationResponse(c *gin.Context, v url.Values) {
	c.JSON(http.StatusOK, gin.H{"error": v})
}

func (s *server) unknownError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}
