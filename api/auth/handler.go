package auth

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPatch, http.MethodPut:
		w.WriteHeader(404)
		fmt.Fprintf(w, "Route not found")
	case http.MethodPost:
		HandleLogin(w, r)
	case http.MethodDelete:
		HandleLogout(w, r)
	}
}
