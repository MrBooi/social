package main

import (
	"log"

	"github.com/mrbooi/social/internal/env"
	store "github.com/mrbooi/social/internal/store/storage"
)

func main() {
	cfg := Config{
		address: env.GetString("ADDRESS", ":8080"),
	}

	storage := store.NewStorage(nil)
	app := &application{
		config: cfg,
		Store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
