package handlers

import (
	"net/http"
	"testing"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/carlosmaniero/budgetgo/usecases"
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
}
