package main

import (
	. "github.com/carlosmaniero/budgetgo/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck)
	r.HandleFunc("/entries/", EntryCreateHandler)

	http.ListenAndServe(":3333", r)
}
