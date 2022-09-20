package main

import (
	"github.com/alexhiggins/go-api/internal/command"
	"github.com/alexhiggins/go-api/internal/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

type server struct {
	config   config
	queries  queries
	commands commands
	logger   *zap.Logger
	router   *gin.Engine
}

type commands struct {
	createAuthor command.CreateAuthorCommand
}

type queries struct {
	fetchAuthor     query.FetchAuthorQuery
	fetchAllAuthors query.FetchAllAuthorsQuery
}

func (s *server) run(port string) error {
	r := s.routes()

	if err := r.Run(port); err != nil {
		s.logger.Error("there was an error calling Run on router: %v", zap.Error(err))
		return err
	}

	return nil
}

func (s *server) statusOk(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}

func (s *server) statusNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}

func (s *server) statusCreated(c *gin.Context, body interface{}) {
	c.JSON(http.StatusCreated, body)
}

func (s *server) statusUnprocessable(c *gin.Context, v url.Values) {
	c.JSON(http.StatusUnprocessableEntity, v)
}

func (s *server) statusBadRequestError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}
