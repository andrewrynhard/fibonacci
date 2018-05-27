// Package ui provides the user interface.
package ui

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

// Serve starts the UI server.
func Serve(port int) error {
	r := mux.NewRouter()

	r.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join("swagger", r.URL.Path[1:]))
	})

	log.Printf("Serving Swagger UI at :%d", port)
	r.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("swagger/ui"))))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))

	return nil
}
