package main

import (
	"log"

	"github.com/mrbooi/social/internal/env"
)

func main() {
	cfg := Config{
		address: env.GetString("ADDRESS", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
