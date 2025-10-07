package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mrbooi/social/internal/auth"
	"github.com/mrbooi/social/internal/db"
	"github.com/mrbooi/social/internal/env"
	"github.com/mrbooi/social/internal/mailer"
	"github.com/mrbooi/social/internal/store/cache"
	store "github.com/mrbooi/social/internal/store/storage"
	"go.uber.org/zap"
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
		address:     env.GetString("ADDRESS", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		env:         env.GetString("ENV", "development"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		db: dbConfig{
			address:      env.GetString("DB_ADDRESS", "postgresql://admin:socialpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxLifeTime:  env.GetString("DB_MAX_LIFE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDRESS", "localhost:6379"),
			pass:    env.GetString("REDIS_PASSWORD", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", "admin"),
				pass: env.GetString("AUTH_BASIC_PASS", "admin"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "social",
			},
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	appDb, err := db.New(cfg.db.address,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxLifeTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer appDb.Close()
	// Mailer
	//mailer := mailer.NewSendgrid(cfg.mail.sendGrid.apiKey, cfg.mail.fromEmail)
	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail)
	if err != nil {
		logger.Fatal(err)
	}

	// Authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.iss,
	)

	logger.Info("database connection pool established")

	// cache
	var redisDb *redis.Client
	if cfg.redisCfg.enabled {
		redisDb = cache.NewRedisClient(
			cfg.redisCfg.addr,
			cfg.redisCfg.pass,
			cfg.redisCfg.db,
		)
		logger.Info("redis database cache connection pool established")
	}

	storage := store.NewStorage(appDb)
	cacheStorage := cache.NewRedisStorage(redisDb)

	app := &application{
		config:        cfg,
		logger:        logger,
		Store:         storage,
		cacheStorage:  cacheStorage,
		mailer:        mailtrap,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
