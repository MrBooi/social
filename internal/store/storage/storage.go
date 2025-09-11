package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, Post *Post) error
		GetByID(ctx context.Context, ID int64) (*Post, error)
	}
	users interface {
		Create(ctx context.Context, user *User) error
	}
	Comments interface {
		GetByPostID(ctx context.Context, PostID int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
