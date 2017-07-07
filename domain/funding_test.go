package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Validating a Founding", t, func() {
		Convey("Given I've a Founding with Name, Limit and Closing Date", func() {
			funding := Funding{
				Name:       "Bank account",
				Limit:      1000,
				Amount:     0,
				ClosingDay: 1,
			}

			Convey("When I validate it", func() {
				errs := funding.Validate()

				Convey("Then the validation passes", func() {
					So(errs, ShouldBeNil)
				})
			})
		})

		Convey("Given I've a Fouding without name and with negative Limit and Closing Date values", func() {
			funding := Funding{
				Name:       "",
				Limit:      -1000,
				Amount:     0,
				ClosingDay: -1,
			}

			Convey("When I validate it", func() {
				errs := funding.Validate()

				Convey("Then the founding is invalid", func() {
					So(errs, ShouldNotBeNil)
				})

				Convey("And have three errors", func() {
					So(len(errs), ShouldEqual, 3)
				})

				Convey("And I can see that the name can't be empty", func() {
					shouldHaveErrorIn(errs, "Name", "The \"Name\" field can't be empty")
				})

				Convey("And I can see that the limit can't be negative", func() {
					shouldHaveErrorIn(errs, "Limit", "The \"Limit\" field can't be negative")
				})

				Convey("And I can see that the closing date should be greater than zero", func() {
					shouldHaveErrorIn(errs, "ClosingDay", "The \"ClosingDay\" field should be greater than zero")
				})
			})
		})
	})
}

func shouldHaveErrorIn(errs []error, field string, message string) {
	founded := false
	var foundedError *ValidationError

	for _, err := range errs {
		validationError := err.(*ValidationError)
		if validationError.Field == field {
			founded = true
			foundedError = validationError
			break
		}
	}

	Convey("The "+field+" field has an error", func() {
		So(founded, ShouldBeTrue)
	})
	Convey("And the message error: "+message, func() {
		So(foundedError.Error(), ShouldEqual, message)
	})
}
