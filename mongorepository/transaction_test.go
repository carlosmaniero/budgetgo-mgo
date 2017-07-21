package mongorepository

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

var transaction = domain.Transaction{
	Description: "My transaction",
	Amount:      10,
	Date:        time.Now(),
	Funding: &domain.Funding{
		ID:         bson.NewObjectId().Hex(),
		Amount:     100,
		ClosingDay: 1,
		Limit:      10,
		Name:       "Bank account",
	},
}

func TestCaseTransaction(t *testing.T) {
	Convey("Scenario: Registering a transaction", t, func() {
		repository := NewMongoTransactionRepository(db.C("test-transaction-on-create"))
		Convey("Given I've a transaction", func() {
			Convey("When I insert it", func() {
				id := repository.Store(&transaction)

				Convey("Then I receive an ObjectId as string", func() {
					So(bson.IsObjectIdHex(id), ShouldBeTrue)
				})
			})
		})
	})
	Convey("Scenario: Get a created transaction", t, func() {
		repository := NewMongoTransactionRepository(db.C("test-transaction-on-retrieve"))

		Convey("Given I've a created transaction", func() {
			id := repository.Store(&transaction)

			Convey("When I get it", func() {
				transactionReceived := repository.FindByID(id)

				Convey("Then I've the same transaction", func() {
					So(transactionReceived.ID, ShouldEqual, id)
					So(transactionReceived.Amount, ShouldEqual, transaction.Amount)
					So(transactionReceived.Funding.ID, ShouldEqual, transaction.Funding.ID)
					So(transactionReceived.Description, ShouldEqual, transaction.Description)
					So(transactionReceived.Date.Unix(), ShouldEqual, transaction.Date.Unix())
				})
			})
		})
	})
}
