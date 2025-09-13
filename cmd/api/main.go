package main

import (
	"log"

	"github.com/mrbooi/social/internal/db"
	"github.com/mrbooi/social/internal/env"
	store "github.com/mrbooi/social/internal/store/storage"
)

const version = "0.0.1"

//	@title			Social GO API
//	@description	API for GO Social, a social network for gohpers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := Config{
		address: env.GetString("ADDRESS", ":8080"),
		apiURL:  env.GetString("EXTERNAL_URL", "localhost:8080"),
		env:     env.GetString("ENV", "development"),
		db: dbConfig{
			address:      env.GetString("DB_ADDRESS", "postgresql://admin:socialpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxLifeTime:  env.GetString("DB_MAX_LIFE_TIME", "15m"),
		},
	}

	appDb, err := db.New(cfg.db.address,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxLifeTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer appDb.Close()

	log.Println("database connection pool established")

	storage := store.NewStorage(appDb)
	app := &application{
		config: cfg,
		Store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
