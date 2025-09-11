package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound         = errors.New("resource not found")
	QueryTmeOutDuration = 5 * time.Second
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, Post *Post) error
		GetByID(ctx context.Context, ID int64) (*Post, error)
		Delete(ctx context.Context, postID int64) error
		Update(ctx context.Context, Post *Post) error
	}
	users interface {
		Create(ctx context.Context, user *User) error
	}
	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
