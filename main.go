package main

import (
	"github.com/carlosmaniero/budgetgo/interfaces/application"
	. "github.com/carlosmaniero/budgetgo/interfaces/handlers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	app := application.Init()
	handlers := Handlers{Application: app}
	router := httprouter.New()

	router.POST("/transaction", handlers.TransactionCreate)
	log.Fatal(http.ListenAndServe(":8123", router))
}
