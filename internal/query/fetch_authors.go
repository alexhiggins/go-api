package query

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/repository"
)

type FetchAllAuthorsQuery struct {
	persistence repository.AuthorRepository
}

func (f *FetchAllAuthorsQuery) Handle(ctx context.Context) ([]data.Author, error) {
	return f.persistence.All(ctx)
}

func NewFetchAllAuthorsQuery(persistence repository.AuthorRepository) FetchAllAuthorsQuery {
	return FetchAllAuthorsQuery{persistence: persistence}
}
