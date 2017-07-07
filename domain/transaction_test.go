package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Scenario: Validating an Transaction", t, func() {
		Convey("Given I've a Transaction with description, amount different of zero and a valid Funding", func() {
			transaction := Transaction{
				Description: "4 beers",
				Amount:      10.5,
				Funding: Funding{
					Name:       "Bank account",
					Limit:      1000,
					Amount:     0,
					ClosingDay: 1,
				},
			}

			Convey("When I try to validate the transaction", func() {
				errors := transaction.Validate()

				Convey("Then the Transaction is valid", func() {
					So(errors, ShouldBeNil)
				})
			})
		})
		Convey("Given I've a Transaction with no description, amount equal zero and a invalid Funding", func() {
			transaction := Transaction{
				Description: "",
				Amount:      0,
				Funding:     Funding{},
			}

			Convey("When I try to validate the transaction", func() {
				errors := transaction.Validate()

				Convey("Then the Transaction isn't valid", func() {
					So(errors, ShouldNotBeNil)
				})

				Convey("And I can see that the Description can't be empty", func() {
					shouldHaveErrorIn(errors, "Description", "The \"Description\" field can't be empty")
				})

				Convey("And I can see that the Amount can't be zero", func() {
					shouldHaveErrorIn(errors, "Amount", "The \"Amount\" field can't be equal zero")
				})

				Convey("And I can see that the Funding is invalid", func() {
					shouldHaveErrorIn(errors, "Funding", "The \"Funding\" field isn't valid")
				})
			})
		})
	})
}
