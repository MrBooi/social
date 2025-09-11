package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	store "github.com/mrbooi/social/internal/store/storage"
)

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "userID")
	userId, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.Store.Users.GetByID(ctx, userId)

	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	if err := jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}

}
