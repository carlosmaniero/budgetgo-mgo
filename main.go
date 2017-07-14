package main

import (
	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/carlosmaniero/budgetgo/interfaces/handlers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	app := application.New()
	appHandlers := handlers.Handlers{Application: app}
	router := httprouter.New()

	router.POST("/transaction", appHandlers.TransactionCreate)
	router.GET("/funding/:id", appHandlers.FundingRetrieve)
	router.POST("/funding", appHandlers.FundingCreate)
	log.Fatal(http.ListenAndServe(":8123", router))
}
