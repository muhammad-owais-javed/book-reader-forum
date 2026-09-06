package handlers

import (
	"context"
	"net/http"
	"time"
)

// wraps a timeout condition around the handler
func withTimeout(w http.ResponseWriter, r *http.Request, handler http.HandlerFunc) {

	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	handler(w, r.WithContext(ctx)) // calling the handler that was passed, but with context
}
