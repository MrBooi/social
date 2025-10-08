# Social

A backend social networking API written in Go using [Chi](https://github.com/go-chi/chi).  
This service provides endpoints for user authentication, posting content, following users, and viewing personalized feeds.

---

## 🚀 Features

- RESTful API with versioning (`/v1`)
- JWT-based authentication and email activation
- CRUD operations for posts
- Follow/unfollow functionality
- Personalized user feed
- Swagger API documentation
- Rate limiting, timeouts, and structured logging
- Environment-based configuration
- Docker and Makefile support for easy development

---

## 🗂️ Project Structure

```
.
├── cmd/api/               → Application entrypoint
├── internal/              → Core business logic and services
│   ├── auth/              → JWT & Basic authentication
│   ├── mailer/            → Email client integrations
│   ├── ratelimiter/       → Request limiting middleware
│   ├── store/             → Storage and caching logic
│   └── env/               → Environment config utilities
├── docs/                  → Swagger docs (auto-generated)
├── Makefile               → Developer commands
├── docker-compose.yml     → Dev environment setup
├── go.mod / go.sum
└── .github/workflows/     → CI/CD configuration
```

---

## ⚙️ Requirements

- Go **1.22+**
- PostgreSQL
- Redis (optional)
- Docker & Docker Compose
- Make

---

## 🏁 Quick Start

1. **Clone the repository**

   ```bash
   git clone https://github.com/MrBooi/social.git
   cd social
   ```

2. **Setup environment**

   ```bash
   cp .envrc.example .envrc
   # Edit environment variables as needed
   ```

3. **Start services**

   ```bash
   docker-compose up --build
   ```

4. **Run the API**

   ```bash
   make run
   ```

5. **Access API documentation**

   Open in browser:

   ```
   http://localhost:8080/v1/swagger/index.html
   ```

---

## 🧩 API Endpoints

### Health & Debug

| Method | Path             | Description           | Auth       |
|--------|-----------------|---------------------|------------|
| GET    | /v1/health      | Health check         | None       |
| GET    | /v1/debug/vars  | Application metrics  | Basic Auth |

### Authentication

| Method | Path                          | Description         | Auth   |
|--------|-------------------------------|-------------------|--------|
| POST   | /v1/authentication/user       | Register a new user| Public |
| POST   | /v1/authentication/token      | Obtain JWT token   | Public |

### Users

| Method | Path                       | Description                     | Auth   |
|--------|----------------------------|---------------------------------|--------|
| PUT    | /v1/users/activate/{token} | Activate user account           | Public |
| GET    | /v1/users/{userID}         | Fetch user profile              | Bearer |
| PUT    | /v1/users/{userID}/follow  | Follow a user                   | Bearer |
| PUT    | /v1/users/{userID}/unfollow| Unfollow a user                 | Bearer |
| GET    | /v1/users/feed             | Get authenticated user’s feed   | Bearer |

### Posts

| Method | Path               | Description                       | Auth   |
|--------|------------------|-----------------------------------|--------|
| POST   | /v1/posts/        | Create a new post                 | Bearer |
| GET    | /v1/posts/{postID}/| Get post by ID                    | Bearer |
| PATCH  | /v1/posts/{postID}/| Update post (moderator or owner) | Bearer |
| DELETE | /v1/posts/{postID}/| Delete post (admin)               | Bearer |

### Swagger Docs

| Method | Path                     | Description |
|--------|-------------------------|-------------|
| GET    | /v1/swagger/index.html  | Swagger UI  |
| GET    | /v1/swagger/doc.json    | API definition |

---

## 🧰 Makefile Usage

| Command | Description |
|---------|-------------|
| make run         | Runs the API locally using `go run ./cmd/api` |
| make build       | Builds the API binary (output to `/bin` or `./social`) |
| make test        | Runs all unit and integration tests |
| make lint        | Runs static analysis tools like `golangci-lint` or `staticcheck` |
| make tidy        | Cleans and tidies `go.mod` and `go.sum` |
| make docker-up   | Starts Docker containers (DB, Redis, etc.) via `docker-compose up` |
| make docker-down | Stops and removes Docker containers |
| make migrate     | Runs database migrations (if supported) |
| make audit       | Runs code security / quality checks in CI |
| make fmt         | Formats code using `go fmt ./...` |
| make clean       | Removes binaries and temporary build files |

Example usage:

```bash
make docker-up     # Start Postgres and Redis
make run           # Run API
make test          # Run all tests
```

---

## ⚙️ Configuration

| Variable | Description | Default |
|----------|-------------|---------|
| APP_ENV               | Application environment | development |
| APP_ADDRESS           | HTTP server address      | :8080 |
| DB_DSN                | PostgreSQL DSN           | — |
| REDIS_ADDR            | Redis address            | — |
| JWT_SECRET            | Secret for JWT tokens    | — |
| CORS_ALLOWED_ORIGIN   | CORS origin for frontend | http://localhost:5174 |
| RATE_LIMIT_ENABLED    | Enable rate limiting     | false |

---

## 🧪 Testing

Run tests:

```bash
make test
```

Or manually:

```bash
go test ./... -v
```

Example endpoint test:

```bash
curl -X POST http://localhost:8080/v1/authentication/user \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

---

## 🐳 Deployment

**Build Docker image:**

```bash
docker build -t social-api .
docker run -p 8080:8080 social-api
```

**Or via Docker Compose:**

```bash
docker-compose up --build
```

---

## 🛣️ Roadmap

- [ ] Add comments and likes
- [ ] Media uploads (images, videos)
- [ ] Real-time updates (WebSockets)
- [ ] User notifications
- [ ] Role-based access (admin/moderator)
- [ ] Improved Swagger documentation

---

## 👥 Author

**Ayabonga Booi**  
GitHub: [@MrBooi](https://github.com/MrBooi)

---

