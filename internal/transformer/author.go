package transformer

import "github.com/alexhiggins/go-api/internal/data"

type Author struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type ShowAuthorResponse struct {
	Data Author `json:"data"`
}

type ShowAllAuthorsResponse struct {
	Data []Author `json:"data"`
}

func ShowAuthor(author data.Author) ShowAuthorResponse {
	return ShowAuthorResponse{
		Data: Author{
			Id:   author.ID,
			Name: author.Name,
			Bio:  author.Bio.String,
		},
	}
}

func ShowAllAuthors(authors []data.Author) ShowAllAuthorsResponse {
	r := ShowAllAuthorsResponse{
		Data: []Author{},
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
