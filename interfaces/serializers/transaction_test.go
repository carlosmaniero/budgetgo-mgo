package serializers

import (
	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"time"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Unserializing one transaction", t, func() {
		serializer := TransactionResponseSerializer{}
		Convey("Given I've a valid transaction json representation", func() {
			jsonTransaction := strings.NewReader("{\"description\": \"4 beers\", \"amount\": 10.50}")

			Convey("Then I can unserialize it", func() {
				err := serializer.Unserialize(jsonTransaction)
				So(err, ShouldBeNil)

				Convey("And see the json data inside the transaction", func() {
					So(serializer.Description, ShouldEqual, "4 beers")
					So(serializer.Amount, ShouldEqual, 10.50)
				})
			})
		})
		Convey("Given I've a transaction with amount as string", func() {
			jsonTransaction := strings.NewReader("{\"description\": \"4 beers\", \"amount\": \"10.50\"}")

			Convey("Then I can't unserialize it", func() {
				err := serializer.Unserialize(jsonTransaction)
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Scenario: Serializing on transaction", t, func() {
		Convey("Given I've a transaction", func() {
			transaction := domain.Transaction{
				ID:          "my-id",
				Description: "5 beers",
				Amount:      22.90,
				Date:        time.Now(),
			}

			Convey("When I serialize it", func() {
				serializer := TransactionResponseSerializer{}
				serializer.Loads(&transaction)
				data := string(serializer.Serialize())

				Convey("Then I can see the data serialized", func() {
					strDate := transaction.Date.Format(time.RFC3339Nano)
					So(data, ShouldEqual, "{\"id\":\"my-id\",\"description\":\"5 beers\",\"amount\":22.9,\"date\":\""+strDate+"\"}")
				})
			})
		})
	})
}
