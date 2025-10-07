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

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	err = writeJSONError(w, http.StatusUnauthorized, "unauthorized")
	if err != nil {
		return
	}
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	err = writeJSONError(w, http.StatusUnauthorized, "unauthorized")
	if err != nil {
		return
	}
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	err := writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
	if err != nil {
		return
	}
}
