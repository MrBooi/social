package main

import (
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("internal server error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem.")
	if err != nil {
		return
	}

}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	err := writeJSONError(w, http.StatusForbidden, "forbidden")
	if err != nil {
		return
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("bad request error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusBadRequest, err.Error())
	if err != nil {
		return
	}
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("conflict error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusConflict, err.Error())
	if err != nil {
		return
	}
}
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("not found error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusNotFound, "the requested resource could not be found.")
	if err != nil {
		return
	}
}
