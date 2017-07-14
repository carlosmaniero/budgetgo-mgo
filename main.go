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
	router.GET("/funding/:id", handlers.FundingRetrieve)
	router.POST("/funding", handlers.FundingCreate)
	log.Fatal(http.ListenAndServe(":8123", router))
}
