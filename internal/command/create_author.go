package command

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
	"github.com/alexhiggins/go-api/internal/repository"
)

type CreateAuthorCommand struct {
	persistence repository.AuthorRepository
}

func (c *CreateAuthorCommand) Handle(ctx context.Context, payload data.CreateAuthorParams) (data.Author, error) {
	return c.persistence.Create(ctx, payload)
}

func NewCreateAuthorCommand(persistence repository.AuthorRepository) CreateAuthorCommand {
	return CreateAuthorCommand{persistence: persistence}
}
