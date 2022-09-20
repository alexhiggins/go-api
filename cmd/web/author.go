package main

import (
	"database/sql"
	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type Author struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type AuthorResponse struct {
	Data Author `json:"data"`
}

type AuthorsResponse struct {
	Data []Author `json:"data"`
}

func (s *server) ShowAuthorsHandler(c *gin.Context) {
	authors, err := s.queries.fetchAllAuthors.Handle(c)

	if err != nil {
		s.unknownError(c, "unable to fetch authors")
		s.logger.Error("unable to fetch authors: %v", zap.Error(err))
		return
	}

	r := AuthorsResponse{Data: []Author{}}
	for _, a := range authors {
		r.Data = append(r.Data, Author{
			Id:   a.ID,
			Name: a.Name,
			Bio:  a.Bio.String,
		})
	}

	s.okResponse(c, r)
}

func (s *server) GetAuthorHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	author, err := s.queries.fetchAuthor.Handle(c, id)

	if err != nil {
		s.notFoundResponse(c, "unable to find author")
		return
	}

	s.okResponse(c, AuthorResponse{Data: Author{
		Id:   author.ID,
		Name: author.Name,
		Bio:  author.Bio.String,
	}})
}

func (s *server) CreateAuthorsHandler(c *gin.Context) {
	v, author := validation.CreateAuthorRequest(c.Request)
	if len(v) > 0 {
		s.failedValidationResponse(c, v)
		return
	}

	createdAuthor, err := s.commands.createAuthor.Handle(c, data.CreateAuthorParams{
		Name: author.Name,
		Bio:  sql.NullString{String: author.Bio, Valid: true},
	})

	if err != nil {
		s.unknownError(c, "unable to create author")
		s.logger.Error("unable to create author: %v", zap.Error(err))
		return
	}

	s.createdResponse(c, AuthorResponse{Data: Author{
		Id:   createdAuthor.ID,
		Name: createdAuthor.Name,
		Bio:  createdAuthor.Bio.String,
	}})
}
