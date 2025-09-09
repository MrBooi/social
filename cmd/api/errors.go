package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem.")
	if err != nil {
		return
	}
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusBadRequest, err.Error())
	if err != nil {
		return
	}
}
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error: %s path: %s error:%s", r.Method, r.URL.Path, err.Error())
	err = writeJSONError(w, http.StatusNotFound, "the requested resource could not be found.")
	if err != nil {
		return
	}
}
