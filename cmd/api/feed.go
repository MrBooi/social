package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// pagination,filters

	ctx := r.Context()
	user := getUserFromContext(r)

	feed, err := app.Store.Posts.GetUserFeed(ctx, user.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
