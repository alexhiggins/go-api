package validation

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"net/url"
)

type CreateAuthorRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func Validate(r *http.Request) (url.Values, CreateAuthorRequest) {
	var a CreateAuthorRequest
	rules := govalidator.MapData{
		"name": []string{"required", "between:3,10"},
		"bio":  []string{"required", "min:50", "max:200"},
	}

	v := govalidator.New(govalidator.Options{
		Request: r,
		Data:    &a,
		Rules:   rules,
	})

	return v.ValidateJSON(), a
}
