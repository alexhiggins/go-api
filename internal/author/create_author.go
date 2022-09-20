package author

import (
	"context"
	"github.com/alexhiggins/go-api/internal/data"
)

type CreateAuthorCommand struct {
	persistence Repository
}

func (c *CreateAuthorCommand) Handle(ctx context.Context, payload data.CreateAuthorParams) (data.Author, error) {
	return c.persistence.Create(ctx, payload)
}

func NewCreateAuthorCommand(persistence Repository) CreateAuthorCommand {
	return CreateAuthorCommand{persistence: persistence}
}
