package utils

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/edward-/four-in-a-row-game/internal/app"
)

func ExecuteRequest(ctx context.Context, req *http.Request) *httptest.ResponseRecorder {
	router := app.Bootstrap(ctx)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)

	return r
}
