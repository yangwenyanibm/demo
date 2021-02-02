package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello {%s}!", vars["name"])
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}