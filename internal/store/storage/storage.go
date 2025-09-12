package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound         = errors.New("resource not found")
	ErrConflict         = errors.New("resource conflict")
	QueryTmeOutDuration = 5 * time.Second
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, Post *Post) error
		GetByID(ctx context.Context, ID int64) (*Post, error)
		Delete(ctx context.Context, postID int64) error
		Update(ctx context.Context, Post *Post) error
	}
	Users interface {
		Create(ctx context.Context, user *User) error
		GetByID(ctx context.Context, userId int64) (*User, error)
	}
	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
		Create(ctx context.Context, comment *Comment) error
	}
	Followers interface {
		Follow(ctx context.Context, followerID, userID int64) error
		Unfollow(ctx context.Context, followerID, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db: db},
	}
}
