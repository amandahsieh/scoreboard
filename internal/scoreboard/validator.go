package scoreboard

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	validate  = validator.New()
	nameRegex = regexp.MustCompile(`^[a-zA-Z0-9 _-]+$`)
)

type Request struct {
	Name string `json:"name" validate:"required,max=255,scoreboard_name"`
}

func init() {
	err := validate.RegisterValidation("scoreboard_name", func(fl validator.FieldLevel) bool {
		return nameRegex.MatchString(fl.Field().String())
	})
	if err != nil {
		return
	}
}

func ValidateScoreboardRequest(req *Request) error {
	return validate.Struct(req)
}
