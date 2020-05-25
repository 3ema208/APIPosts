package store

import "github.com/3ema208/APIPosts/internal/app/model"

// CommentStore ...
type CommentStore struct {
	store *Store
}

// Create ...
func (p *CommentStore) Create(comm *model.Comment) (*model.Comment, error) {
	err := p.store.db.QueryRow(
		"INSERT INTO comment (author, content, post_id) VALUES ($1, $2, $3) RETURNING id",
		comm.Author,
		comm.Content,
		comm.Post.ID,
	).Scan(&comm.ID)
	if err != nil {
		return nil, err

	}
	return comm, nil
}
