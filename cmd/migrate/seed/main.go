package main

import (
	"database/sql"
	"log"

	"github.com/mrbooi/social/internal/db"
	"github.com/mrbooi/social/internal/env"
	store "github.com/mrbooi/social/internal/store/storage"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	newStore := store.NewStorage(conn)

	db.Seed(newStore, conn)
}
