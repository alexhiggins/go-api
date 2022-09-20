package validation

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type NewAuthor struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func CreateAuthorRequest(r *http.Request) (url.Values, NewAuthor) {
	var author NewAuthor
	rules := govalidator.MapData{
		"name": []string{"required", "between:3,10"},
		"bio":  []string{"required", "min:50", "max:200"},
	}

	v := govalidator.New(govalidator.Options{
		Request: r,
		Data:    &author,
		Rules:   rules,
	})

	return v.ValidateJSON(), author
}
