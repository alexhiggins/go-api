package repository

import (
	"context"
	"database/sql"
	"github.com/alexhiggins/go-api/internal/data"
)

type AuthorRepository interface {
	All(ctx context.Context) ([]data.Author, error)
	WhereId(ctx context.Context, id int64) (data.Author, error)
	Create(ctx context.Context, params data.CreateAuthorParams) (data.Author, error)
}

type AuthorDbRepository struct {
	query *data.Queries
}

func (d *AuthorDbRepository) All(ctx context.Context) ([]data.Author, error) {
	return d.query.ListAuthors(ctx)
}

func (d *AuthorDbRepository) WhereId(ctx context.Context, id int64) (data.Author, error) {
	return d.query.GetAuthor(ctx, id)
}

func (d *AuthorDbRepository) Create(ctx context.Context, params data.CreateAuthorParams) (data.Author, error) {
	return d.query.CreateAuthor(ctx, params)
}

func NewAuthorRepository(db *sql.DB) *AuthorDbRepository {
	return &AuthorDbRepository{query: data.New(db)}
}
