package main

import (
	"log"
	"net/http"

	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/carlosmaniero/budgetgo/interfaces/handlers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	app := application.NewWithMongo(session)
	appHandlers := handlers.Handlers{Application: app}
	router := httprouter.New()

	router.POST("/transaction", appHandlers.TransactionCreate)
	router.GET("/transaction/:id", appHandlers.TransactionRetrieve)
	router.GET("/funding/:id", appHandlers.FundingRetrieve)
	router.POST("/funding", appHandlers.FundingCreate)
	log.Fatal(http.ListenAndServe(":8123", router))
}
