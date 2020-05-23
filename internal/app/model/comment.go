package model

// Comment ..
type Comment struct {
	ID int 
	Author string
	Content string
	CreatedTime string
	Post *Post
}
