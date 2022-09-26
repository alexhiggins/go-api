package transformer

import (
	"github.com/alexhiggins/go-api/internal/data"
)

type Error struct {
	Error string `json:"error" example:"unable to perform action"`
}

type ShowHealthResponse struct {
	Status string `json:"status" example:"OK"`
}

type ItemResponse struct {
	Data interface{} `json:"data"`
}

type ListResponse struct {
	Data []interface{} `json:"data"`
}

type Author struct {
	Id   int64  `json:"id" example:"1"`
	Name string `json:"name" example:"john doe"`
	Bio  string `json:"bio" example:"john doe's biography description'"`
}

func ShowHealth(status string) ShowHealthResponse {
	return ShowHealthResponse{
		Status: status,
	}
}

func ShowAuthor(author data.Author) ItemResponse {
	return ItemResponse{
		Data: Author{
			Id:   author.ID,
			Name: author.Name,
			Bio:  author.Bio.String,
		},
	}
}

func ShowAllAuthors(authors []data.Author) ListResponse {
	r := ListResponse{
		Data: []interface{}{},
	}

	for _, a := range authors {
		r.Data = append(r.Data, Author{
			Id:   a.ID,
			Name: a.Name,
			Bio:  a.Bio.String,
		})
	}

	return r
}
