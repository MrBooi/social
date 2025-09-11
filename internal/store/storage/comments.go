package store

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	UserID    int64     `json:"user_id"`
	PostID    int64     `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByPostID(ctx context.Context, PostID int64) ([]Comment, error) {
	var query = ` SELECT c.id, c.post_id, c.user_id,c.content,c.created_at,users.username,users.id FROM comments c
     join users on users.id = c.user_id
    WHERE c.post_id = $1
 	order by c.created_at desc
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTmeOutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, PostID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var comments []Comment
	for rows.Next() {
		var c Comment
		c.User = User{}
		err := rows.Scan(
			&c.ID,
			c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.Username,
			&c.User.ID,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
