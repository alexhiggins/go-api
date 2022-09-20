package main

import (
	"database/sql"
	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/presenter"
	"github.com/alexhiggins/go-api/internal/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func (s *server) ShowAuthorsHandler(c *gin.Context) {
	authors, err := s.queries.fetchAllAuthors.Handle(c)

	if err != nil {
		s.statusBadRequestError(c, "unable to fetch authors")
		s.logger.Error("unable to fetch authors: %v", zap.Error(err))
		return
	}

	s.statusOk(c, presenter.ShowAllAuthors(authors))
}

func (s *server) GetAuthorHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	a, err := s.queries.fetchAuthor.Handle(c, id)

	if err != nil {
		s.statusNotFound(c, "unable to find author")
		return
	}

	s.statusOk(c, presenter.ShowAuthor(a))
}

func (s *server) CreateAuthorHandler(c *gin.Context) {
	v, a := validation.Validate(c.Request)
	if len(v) > 0 {
		s.statusUnprocessable(c, v)
		return
	}

	newAuthor, err := s.commands.createAuthor.Handle(c, data.CreateAuthorParams{
		Name: a.Name,
		Bio:  sql.NullString{String: a.Bio, Valid: true},
	})

	if err != nil {
		s.statusBadRequestError(c, "unable to create author")
		s.logger.Error("unable to create author: %v", zap.Error(err))
		return
	}

	s.statusCreated(c, presenter.ShowAuthor(newAuthor))
}
