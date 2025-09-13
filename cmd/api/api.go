package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrbooi/social/docs" // generated doc

	store "github.com/mrbooi/social/internal/store/storage" // swagger embed files
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type application struct {
	config Config
	Store  store.Storage
}

type Config struct {
	address string
	db      dbConfig
	env     string
	apiURL  string
}

type dbConfig struct {
	address      string
	maxOpenConns int
	maxIdleConns int
	maxLifeTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		// http://localhost:8080/v1/swagger/index.html
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.address)
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(docsURL), //The url pointing to API definition
		))
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)

			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)

				r.Get("/", app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/", app.updatePostHandler)
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)
				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})

		})

	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route not found: "+r.URL.Path, http.StatusNotFound)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

	server := &http.Server{
		Addr:         app.config.address,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("API server listening on %s", app.config.address)
	return server.ListenAndServe()
}
