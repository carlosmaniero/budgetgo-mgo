package handlers

import (
	"net/http"
	"testing"

	"github.com/carlosmaniero/budgetgo/interfaces/application"
	"github.com/julienschmidt/httprouter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecError(t *testing.T) {
	Convey("Scenario: Handlering an error", t, func() {
		app := application.Init()
		handlers := Handlers{Application: app}

		Convey("Given I've a handler that raise a panic", func() {
			errorResponse := HandlerResponseMock{}
			request := http.Request{
				Body: NewRequestBodyMock("{}"),
			}

			Convey("When I sent it to the handler", func() {
				handlers.panicHandler(&errorResponse, &request, nil)

				Convey("Then the error was created successly", func() {
					So(errorResponse.ResponseBody, ShouldEqual, "{\"type\":\"server_error\",\"message\":\"An error was occurred check your request body\"}")
					So(errorResponse.StatusCode, ShouldEqual, 500)
				})
			})

			Convey("Given I sent unserializable data", func() {
				errorResponse := HandlerResponseMock{}
				request := http.Request{
					Body: NewRequestBodyMock("{\"amount\": \"this is not a number\"}"),
				}

				Convey("When I sent it to the handler", func() {
					handlers.TransactionCreate(&errorResponse, &request, nil)

					Convey("Then the error was created successly", func() {
						So(errorResponse.ResponseBody, ShouldEqual, "{\"type\":\"parser\",\"message\":\"cannot add string value into field amount of type float64\"}")
						So(errorResponse.StatusCode, ShouldEqual, 400)
					})
				})
			})
			Convey("Given I sent unserializable date", func() {
				errorResponse := HandlerResponseMock{}
				request := http.Request{
					Body: NewRequestBodyMock("{\"date\": \"25/09/1993\"}"),
				}

				Convey("When I sent it to the handler", func() {
					handlers.TransactionCreate(&errorResponse, &request, nil)

					Convey("Then the error was created successly", func() {
						So(errorResponse.ResponseBody, ShouldEqual, "{\"type\":\"parser\",\"message\":\"cannot parse the sent date. Check the date format. Date Formate: 2006-01-02T15:04:05Z07:00 (RFC3339)\"}")
						So(errorResponse.StatusCode, ShouldEqual, 400)
					})
				})
			})
		})
	})
}

func (handlers *Handlers) panicHandler(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer handlers.catchPanics(response)

	panic("Ow! I can't do it.")
}
