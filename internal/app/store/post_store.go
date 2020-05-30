package store

import "github.com/3ema208/APIPosts/internal/app/model"

// PostStore ..
type PostStore struct {
	store *Store
}

// Create Post
func (p *PostStore) Create(postModel *model.Post) (*model.Post, error) {
	err := p.store.db.QueryRow(
		"INSERT INTO posts (title, link, author_name) VALUES ($1, $2, $3) RETURNING id",
		postModel.Title,
		postModel.Link,
		postModel.AuthorName,
	).Scan(&postModel.ID)

	if err != nil {
		return nil, err
	}

	return postModel, nil
}

// Get ...
func (p *PostStore) Get() ([]*model.Post, error) {
	var posts []*model.Post
	rows, err := p.store.db.Query("SELECT id, title, link, author_name, created_time, amount_upvote FROM posts ORDER BY id DESC LIMIT $1", 100)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := &model.Post{}
		errScan := rows.Scan(&p.ID, &p.Title, &p.Link, &p.AuthorName, &p.CreatedTime, &p.AmountUpvote)
		if errScan != nil {
			return nil, errScan
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// FindByID find post by id
func (p *PostStore) FindByID(postID string) (*model.Post, error) {
	post := &model.Post{}
	err := p.store.db.QueryRow("SELECT id, title, link, author_name, created_time, amount_upvote FROM posts WHERE id=$1", postID).Scan(
		&post.ID,
		&post.Title,
		&post.Link,
		&post.AuthorName,
		&post.CreatedTime,
		&post.AmountUpvote,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// Update ...
func (p *PostStore) Update(postID string) (*model.Post, error) {
	return nil, nil
}

// Delete ...
func (p *PostStore) Delete(postID string) (*model.Post, error) {
	return nil, nil
}
