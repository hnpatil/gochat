package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const UserID = "X-User-ID"

func Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if shouldSkip(r) {
			handler.ServeHTTP(w, r)
			return
		}

		userID := r.Header.Get(UserID)
		if userID == "" && !(r.Method == http.MethodGet && r.URL.Path == "/v1/users") {
			http.Error(w, fmt.Sprintf("Bad Request: %s header missing", UserID), http.StatusBadRequest)
			return
		}

		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserID, userID)))
	})
}

func shouldSkip(r *http.Request) bool {
	return strings.HasPrefix(r.URL.Path, "/.well-known")
}
