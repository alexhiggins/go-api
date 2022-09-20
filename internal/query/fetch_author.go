package query

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/repository"
)

type FetchAuthorQuery struct {
	persistence repository.AuthorRepository
}

func (f *FetchAuthorQuery) Handle(ctx context.Context, id int64) (data.Author, error) {
	return f.persistence.WhereId(ctx, id)
}

func NewFetchAuthorQuery(persistence repository.AuthorRepository) FetchAuthorQuery {
	return FetchAuthorQuery{persistence: persistence}
}
