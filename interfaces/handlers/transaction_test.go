package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecTransaction(t *testing.T) {
	Convey("Scenario: Registering an transaction", t, func() {
		app := application.New()
		handlers := Handlers{Application: app}
		transactionResponse := HandlerResponseMock{}

		fundingIteractor := usecases.FundingInteractor{Repository: app.FundingRepository}
		funding := domain.Funding{
			Name:       "Beer account",
			Amount:     1,
			ClosingDay: 2,
			Limit:      3,
		}

		if err := fundingIteractor.Register(&funding); err != nil {
			panic(err)
		}

		Convey("Given I've a valid transaction json representation", func() {
			now := time.Now().Format(time.RFC3339Nano)

			request := http.Request{
				Body: NewRequestBodyMock("{\"description\": \"8 beers\", \"funding_id\": \"" + funding.ID + "\", \"amount\": 10, \"date\": \"" + now + "\"}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.TransactionCreate(&transactionResponse, &request, nil)

				Convey("Then the transaction was created successly", func() {
					So(transactionResponse.ResponseBody, ShouldEqual, "{\"id\":\"1\",\"description\":\"8 beers\",\"amount\":10,\"date\":\""+now+"\",\"funding_id\":\"1\"}")
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
					So(transactionResponse.ResponseBody, ShouldContainSubstring, "the funding was not found")
					So(transactionResponse.StatusCode, ShouldEqual, 400)
				})
			})
		})
	})

	Convey("Scenario: Retrieving a transaction", t, func() {
		app := application.New()
		handlers := Handlers{Application: app}

		Convey("Given I've a created transaction", func() {
			fundingIterator := usecases.FundingInteractor{Repository: app.FundingRepository}
			funding := domain.Funding{
				Name:       "Pet account",
				Amount:     1,
				ClosingDay: 2,
				Limit:      3,
			}

			if err := fundingIterator.Register(&funding); err != nil {
				panic(err)
			}

			iteractor := usecases.TransactionInteractor{Repository: app.TransactionRepository}
			transaction := domain.Transaction{
				Description: "Golpher Shampoo",
				Amount:      25.50,
				Date:        time.Now(),
				Funding:     &funding,
			}

			if err := iteractor.Register(&transaction); err != nil {
				panic(err)
			}

			Convey("When I get it from the transaction entrypoint", func() {
				transactionResponse := HandlerResponseMock{}
				request := http.Request{}
				params := make(httprouter.Params, 0)
				params = append(params, httprouter.Param{Key: "id", Value: transaction.ID})
				handlers.TransactionRetrieve(&transactionResponse, &request, params)

				Convey("Then the transaction is returned", func() {
					So(transactionResponse.ResponseBody, ShouldContainSubstring, "{\"id\":\"1\",\"description\":\"Golpher Shampoo\",\"amount\":25.5,\"date\":\""+transaction.Date.Format(time.RFC3339Nano)+"\"}")
				})
			})

			Convey("Given I've a uncreated transaction", func() {
				Convey("When I get it from the transaction entrypoint", func() {
					transactionResponse := HandlerResponseMock{}
					request := http.Request{}
					params := make(httprouter.Params, 0)
					params = append(params, httprouter.Param{Key: "id", Value: "666"})
					handlers.TransactionRetrieve(&transactionResponse, &request, params)

					Convey("Then I can see that the transaction does not exists", func() {
						So(transactionResponse.ResponseBody, ShouldContainSubstring, "{\"type\":\"not-found\",\"message\":\"the transaction was not found\"}")
						So(transactionResponse.StatusCode, ShouldEqual, 404)
					})
				})
			})
		})
	})
}
