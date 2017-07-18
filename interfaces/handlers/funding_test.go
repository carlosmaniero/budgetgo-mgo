package handlers

import (
	"net/http"
	"testing"

	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/carlosmaniero/budgetgo/usecases"
	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Registering a funding", t, func() {
		app := application.New()
		handlers := Handlers{Application: app}
		fundingResponse := HandlerResponseMock{}

		Convey("Given I've a valid transaction json representation", func() {
			request := http.Request{
				Body: NewRequestBodyMock("{\"name\": \"Bank Account\", \"closing_day\": 1}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.FundingCreate(&fundingResponse, &request, nil)

				Convey("Then the funding was created successly", func() {
					So("{\"id\":\"1\",\"name\":\"Bank Account\",\"limit\":0,\"amount\":0,\"closing_day\":1}", ShouldEqual, fundingResponse.ResponseBody)
					So(fundingResponse.StatusCode, ShouldEqual, 201)
				})
			})
		})

		Convey("Given I've a invalid transaction json representation", func() {
			request := http.Request{
				Body: NewRequestBodyMock("{\"name\": \"\", \"closing_day\": 1}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.FundingCreate(&fundingResponse, &request, nil)

				Convey("Then the funding was not created successly", func() {
					So(fundingResponse.ResponseBody, ShouldContainSubstring, "This entity is not valid")
					So(fundingResponse.StatusCode, ShouldEqual, 400)
				})
			})
		})
	})

	Convey("Scenario: Retrieving a funding", t, func() {
		app := application.New()
		handlers := Handlers{Application: app}

		Convey("Given I've a created funding", func() {
			iteractor := usecases.FundingInteractor{Repository: app.FundingRepository}
			funding := domain.Funding{
				Name:       "Beer account",
				Amount:     1,
				ClosingDay: 2,
				Limit:      3,
			}

			if err := iteractor.Register(&funding); err != nil {
				panic(err)
			}

			Convey("When I get it from the funding entrypoint", func() {
				fundingResponse := HandlerResponseMock{}
				request := http.Request{}
				params := make(httprouter.Params, 0)
				params = append(params, httprouter.Param{Key: "id", Value: funding.ID})
				handlers.FundingRetrieve(&fundingResponse, &request, params)

				Convey("Then the funding is returned", func() {
					So(fundingResponse.ResponseBody, ShouldContainSubstring, "{\"id\":\"1\",\"name\":\"Beer account\",\"limit\":3,\"amount\":1,\"closing_day\":2}")
				})
			})

			Convey("Given I've a uncreated funding", func() {
				Convey("When I get it from the funding entrypoint", func() {
					fundingResponse := HandlerResponseMock{}
					request := http.Request{}
					params := make(httprouter.Params, 0)
					params = append(params, httprouter.Param{Key: "id", Value: "666"})
					handlers.FundingRetrieve(&fundingResponse, &request, params)

					Convey("Then I can see that the funding does not exists", func() {
						So(fundingResponse.ResponseBody, ShouldContainSubstring, "{\"type\":\"not-found\",\"message\":\"the funding was not found\"}")
						So(fundingResponse.StatusCode, ShouldEqual, 404)
					})
				})
			})
		})
	})
}
