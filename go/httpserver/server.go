// Package httpserver: net/http se chhota JSON API, koi framework nahi chahiye.
package httpserver

import (
	"encoding/json"
	"net/http"
)

type health struct {
	Status string `json:"status"` // struct tag: JSON me field ka naam
}

// NewMux alag function me hai taaki test me bina port khole use kar sako
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(health{Status: "ok"})
	})
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("namaste " + r.PathValue("name"))) // Go 1.22+ path params
	})
	return mux
}
