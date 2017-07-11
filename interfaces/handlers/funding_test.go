package handlers

import (
	"net/http"
	"testing"

	"github.com/carlosmaniero/budgetgo/interfaces/application"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Registering an funding", t, func() {
		app := application.Init()
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
					So(fundingResponse.ResponseBody, ShouldContainSubstring, "The funding has validation errors")
					So(fundingResponse.StatusCode, ShouldEqual, 400)
				})
			})
		})
	})
}
