package author

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
)

type FetchAuthorQuery struct {
	persistence Repository
}

func (f *FetchAuthorQuery) Handle(ctx context.Context, id int64) (data.Author, error) {
	return f.persistence.WhereId(ctx, id)
}

func NewFetchAuthorQuery(persistence Repository) FetchAuthorQuery {
	return FetchAuthorQuery{persistence: persistence}
}
