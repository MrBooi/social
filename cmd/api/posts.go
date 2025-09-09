package main

import (
	"net/http"

	store "github.com/mrbooi/social/internal/store/storage"
)

type CreatePostPayload struct {
	//ID      int64    `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
	//UserID  int64    `json:"user_id"`
	Tags []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload

	if err := readJSON(w, *r, &payload); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Content: payload.Content,
		Title:   payload.Title,
		Tags:    payload.Tags,
		// TODO dont forget to change
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.Store.Posts.Create(ctx, post); err != nil {
	_:
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
	_:
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
