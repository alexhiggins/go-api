package main

import (
	"database/sql"
	"strconv"

	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/transformer"
	"github.com/alexhiggins/go-api/internal/validate"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ShowAuthorsHandler godoc
// @Summary Show all authors
// @Tags author
// @Accept */*
// @Produce json
// @Success 201 {object} transformer.ListResponse{data=[]transformer.Author}
// @Failure 400 {object} transformer.Error
// @Router /api/v1/authors [get]
func (s *server) ShowAuthorsHandler(c *gin.Context) {
	authors, err := s.queries.fetchAllAuthors.Handle(c)

	if err != nil {
		s.statusBadRequestError(c, "unable to fetch authors")
		s.logger.Error("unable to fetch authors: %v", zap.Error(err))
		return
	}

	s.statusOk(c, transformer.ShowAllAuthors(authors))
}

// GetAuthorHandler godoc
// @Summary Show an individual author
// @Tags author
// @Accept */*
// @Produce json
// @Param  id path int true "Author ID"
// @Success 200 {object} transformer.ItemResponse{data=transformer.Author}
// @Failure 400 {object} transformer.Error
// @Failure 404 {object} transformer.Error
// @Router /api/v1/authors/{id} [get]
func (s *server) GetAuthorHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	a, err := s.queries.fetchAuthor.Handle(c, id)

	if err != nil {
		s.statusNotFound(c, "unable to find author")
		return
	}

	s.statusOk(c, transformer.ShowAuthor(a))
}

// CreateAuthorHandler godoc
// @Summary Create an author
// @Tags author
// @Accept json
// @Produce json
// @Param account body validate.ValidatedAuthor true  "Create Author"
// @Success 201 {object} transformer.ItemResponse{data=transformer.Author}
// @Failure 400 {object} transformer.Error
// @Router /api/v1/authors [post]
func (s *server) CreateAuthorHandler(c *gin.Context) {
	v, a := validate.NewAuthor(c.Request)
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

	s.statusCreated(c, transformer.ShowAuthor(newAuthor))
}
