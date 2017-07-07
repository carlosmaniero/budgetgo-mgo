package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Scenario: Validating an Transaction", t, func() {
		Convey("Given I've an Transaction with description", func() {
			transaction := Transaction{Description: "4 beers"}

			Convey("And amount different of zero", func() {
				transaction.Amount = 10.5

				Convey("When I try to validate the transaction", func() {
					errors := transaction.Validate()

					Convey("Then the Transaction is valid", func() {
						So(errors, ShouldBeNil)
					})
				})
			})
			Convey("And amount equal zero", func() {
				transaction.Amount = 0

				Convey("When I try to validate the transaction", func() {
					errors := transaction.Validate()

					Convey("Then the Transaction isn't valid", func() {
						So(errors, ShouldNotBeNil)
					})

					Convey("And show me an error", func() {
						So(errors[0].Error(), ShouldEqual, "Amount can't be equal zero")
					})
				})
			})
		})
		Convey("Given I've an Transaction with no description", func() {
			transaction := Transaction{Description: ""}

			Convey("And amount different of zero", func() {
				transaction.Amount = 2

				Convey("When I try to validate the transaction", func() {
					errors := transaction.Validate()

					Convey("Then the Transaction isn't valid", func() {
						So(errors, ShouldNotBeNil)
					})

					Convey("And show me an error", func() {
						So(errors[0].Error(), ShouldEqual, "Description can't be empty")
					})
				})
			})

			Convey("And amount equal zero", func() {
				transaction.Amount = 0

				Convey("When I try to validate the transaction", func() {
					errors := transaction.Validate()

					Convey("Then the Transaction isn't valid", func() {
						So(errors, ShouldNotBeNil)
					})

					Convey("And show me two errors", func() {
						So(len(errors), ShouldEqual, 2)
					})
				})
			})
		})
	})
}
