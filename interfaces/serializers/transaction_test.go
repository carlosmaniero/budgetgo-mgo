package serializers

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Unserializing one transaction", t, func() {
		Convey("Given I've a valid transaction json representation", func() {
			jsonTransaction := strings.NewReader("{\"description\": \"4 beers\", \"amount\": 10.50}")

			Convey("Then I can unserialize it", func() {
				transaction, err := UnserializeTransactionData(jsonTransaction)
				So(err, ShouldBeNil)

				Convey("And see the json data inside the transaction", func() {
					So(transaction.Description, ShouldEqual, "4 beers")
					So(transaction.Amount, ShouldEqual, 10.50)
				})
			})
		})
		Convey("Given I've a transaction with amount as string", func() {
			jsonTransaction := strings.NewReader("{\"description\": \"4 beers\", \"amount\": \"10.50\"}")

			Convey("Then I can't unserialize it", func() {
				_, err := UnserializeTransactionData(jsonTransaction)
				So(err, ShouldNotBeNil)
			})
		})
	})
}