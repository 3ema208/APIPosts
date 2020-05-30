package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Comment ..
type Comment struct {
	ID          int
	Author      string
	Content     string
	CreatedTime string
	Post        *Post
}

// Validate ...
func (c *Comment) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Author, validation.Required, validation.Length(6, 255)),
		validation.Field(&c.Content, validation.Required),
		validation.Field(&c.Post, validation.Required, is.Digit),
	)
}
