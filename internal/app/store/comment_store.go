package store

import "github.com/3ema208/APIPosts/internal/app/model"

// CommentStore ...
type CommentStore struct {
	store *Store
}

// Create ...
func (p *CommentStore) Create(comm *model.Comment) error {
	if errVal := comm.Validate(); errVal != nil {
		return errVal
	}

	err := p.store.db.QueryRow(
		"INSERT INTO comment (author, content, post_id) VALUES ($1, $2, $3) RETURNING id",
		comm.Author,
		comm.Content,
		comm.PostID,
	).Scan(&comm.ID)
	if err != nil {
		return err
	}
	return nil
}

// Get all comment by postID
func (p *CommentStore) Get(postID string) ([]*model.Comment, error) {
	rows, err := p.store.db.Query(
		"SELECT id, author, content, created_time FROM comment WHERE post_id=$1 ORDER BY id DESC",
		postID,
	)
	if err != nil {
		return nil, err
	}
	comments := []*model.Comment{}
	for rows.Next() {
		comment := &model.Comment{}
		rows.Scan(&comment.ID, &comment.Author, &comment.Content, &comment.CreatedTime)
	}
	return comments, nil
}
