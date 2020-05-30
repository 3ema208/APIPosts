package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Post model
type Post struct {
	ID           int
	Title        string
	Link         string
	AuthorName   string
	CreatedTime  string
	AmountUpvote int
}

// Validate fields
func (p *Post) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Title, validation.Required, validation.Length(6, 255)),
		validation.Field(&p.Link, validation.Required, is.URL),
		validation.Field(&p.AuthorName, validation.Required),
	)
}
