package serializers

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"github.com/carlosmaniero/budgetgo/domain"
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

	Convey("Scenario: Serializing on transaction", t, func() {
		Convey("Given I've a transaction", func() {
			transaction := domain.Transaction{
				Description: "5 beers",
				Amount: 22.90,
			}

			Convey("When I serialize it", func() {
				data := string(SerializeTransaction(&transaction))

				Convey("Then I can see the data serialized", func() {
					So(data, ShouldEqual, "{\"description\":\"5 beers\",\"amount\":22.9}")
				})
			})
		})
	})
}