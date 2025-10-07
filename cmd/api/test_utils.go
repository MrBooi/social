package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrbooi/social/internal/auth"
	"github.com/mrbooi/social/internal/store/cache"
	store "github.com/mrbooi/social/internal/store/storage"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg Config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	// Uncomment to enable logs
	// logger := zap.Must(zap.NewProduction()).Sugar()
	mockStore := store.NewMockStore()
	mockCacheStore := cache.NewMockStore()

	testAuth := &auth.TestAuthenticator{}

	// Rate limiter
	//rateLimiter := ratelimiter.NewFixedWindowLimiter(
	//	cfg.rateLimiter.RequestsPerTimeFrame,
	//	cfg.rateLimiter.TimeFrame,
	//)

	return &application{
		logger:        logger,
		Store:         mockStore,
		cacheStorage:  mockCacheStore,
		authenticator: testAuth,
		config:        cfg,
		//rateLimiter:   rateLimiter,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
