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
		GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(ctx context.Context, userId int64) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
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

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
