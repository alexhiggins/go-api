package author

import (
	"context"
	"database/sql"
	"github.com/alexhiggins/go-api/internal/data"
)

type Repository interface {
	All(ctx context.Context) ([]data.Author, error)
	WhereId(ctx context.Context, id int64) (data.Author, error)
	Create(ctx context.Context, params data.CreateAuthorParams) (data.Author, error)
}

type DbRepository struct {
	query *data.Queries
}

func (d *DbRepository) All(ctx context.Context) ([]data.Author, error) {
	return d.query.ListAuthors(ctx)
}

func (d *DbRepository) WhereId(ctx context.Context, id int64) (data.Author, error) {
	return d.query.GetAuthor(ctx, id)
}

func (d *DbRepository) Create(ctx context.Context, params data.CreateAuthorParams) (data.Author, error) {
	return d.query.CreateAuthor(ctx, params)
}

func NewDbRepository(db *sql.DB) *DbRepository {
	return &DbRepository{query: data.New(db)}
}
