package repository

import (
	"context"
	"database/sql"
	"errors"
	"sync"

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

type AuthorInMemoryRepository struct {
	lock    *sync.RWMutex
	authors []data.Author
}

func (d *AuthorInMemoryRepository) All(ctx context.Context) ([]data.Author, error) {
	return d.authors, nil
}

func (d *AuthorInMemoryRepository) WhereId(ctx context.Context, id int64) (data.Author, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	for _, a := range d.authors {
		if a.ID == id {
			return a, nil
		}
	}

	return data.Author{}, errors.New("unable to find author")
}

func (d *AuthorInMemoryRepository) Create(ctx context.Context, params data.CreateAuthorParams) (data.Author, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	author := data.Author{
		ID:   int64(len(d.authors) + 1),
		Name: params.Name,
		Bio:  params.Bio,
	}

	d.authors = append(d.authors, author)

	return author, nil
}

func NewDbAuthorRepository(db *sql.DB) *AuthorDbRepository {
	return &AuthorDbRepository{
		query: data.New(db),
	}
}

func NewInMemoryAuthorRepository() *AuthorInMemoryRepository {
	return &AuthorInMemoryRepository{
		authors: []data.Author{},
		lock:    &sync.RWMutex{},
	}
}
