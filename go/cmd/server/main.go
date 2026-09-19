// Server start karne ke liye: go run ./cmd/server
package main

import (
	"log"
	"net/http"

	"github.com/tanishk/engineering-playbook/go/httpserver"
)

func main() {
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewMux()))
}
