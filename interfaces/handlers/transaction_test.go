package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/carlosmaniero/budgetgo/interfaces/application"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecTransaction(t *testing.T) {
	Convey("Scenario: Registering an transaction", t, func() {
		app := application.Init()
		handlers := Handlers{Application: app}
		transactionResponse := HandlerResponseMock{}

		Convey("Given I've a valid transaction json representation", func() {
			now := time.Now().Format(time.RFC3339Nano)

			request := http.Request{
				Body: NewRequestBodyMock("{\"description\": \"8 beers\", \"amount\": 10, \"date\": \"" + now + "\"}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.TransactionCreate(&transactionResponse, &request, nil)

				Convey("Then the transaction was created successly", func() {
					So(transactionResponse.ResponseBody, ShouldEqual, "{\"id\":\"1\",\"description\":\"8 beers\",\"amount\":10,\"date\":\""+now+"\"}")
					So(transactionResponse.StatusCode, ShouldEqual, 201)
				})
			})
		})

		Convey("Given I've a invalid transaction json representation", func() {
			request := http.Request{
				Body: NewRequestBodyMock("{}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.TransactionCreate(&transactionResponse, &request, nil)

				Convey("Then the transaction was not created successly", func() {
					So(transactionResponse.ResponseBody, ShouldContainSubstring, "The transaction has validation errors")
					So(transactionResponse.StatusCode, ShouldEqual, 400)
				})
			})
		})
	})
}
