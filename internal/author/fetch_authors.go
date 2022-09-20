package author

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
)

type FetchAllAuthorsQuery struct {
	persistence Repository
}

func (f *FetchAllAuthorsQuery) Handle(ctx context.Context) ([]data.Author, error) {
	return f.persistence.All(ctx)
}

func NewFetchAllAuthorsQuery(persistence Repository) FetchAllAuthorsQuery {
	return FetchAllAuthorsQuery{persistence: persistence}
}
