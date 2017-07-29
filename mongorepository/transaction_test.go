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
		PaymentDay: 1,
		Limit:      10,
		Name:       "Bank account",
	},
}

func TestCaseTransaction(t *testing.T) {

	Convey("Scenario: Registering a transaction", t, func() {
		db.DropDatabase()
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
		db.DropDatabase()
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
	Convey("Scenario: Get a list of transactions based on date and funding", t, func() {
		db.DropDatabase()
		repository := NewMongoTransactionRepository(db.C("test-transaction-on-funding-get"))

		Convey("Given I've a List of transactions with different dates", func() {
			initialDate := time.Date(2017, time.June, 10, 0, 0, 0, 0, time.Local)

			storedDates := make(map[string]time.Time)
			dateList := make([]time.Time, 0)

			for day := 45; day > 0; day-- {
				transaction.Date = initialDate.AddDate(0, 0, day)
				id := repository.Store(&transaction)
				storedDates[id] = transaction.Date
				dateList = append(dateList, transaction.Date)
			}

			Convey("When I find another by a nonexistent funding", func() {
				Convey("Then I can't see transactions", func() {
					funding := domain.Funding{
						ID: bson.NewObjectId().Hex(),
					}
					list := repository.FindByFundingAndInterval(&funding, initialDate.AddDate(-100, 0, 0), initialDate.AddDate(100, 0, 0))
					So(list.Next(&domain.Transaction{}), ShouldBeFalse)
				})
			})

			Convey("When I find the list based in a funding and the interval that contains all registered dated", func() {
				list := repository.FindByFundingAndInterval(transaction.Funding, initialDate, initialDate.AddDate(0, 0, 46))
				total := 0

				Convey("Then I can iterate over all dates", func() {
					transaction := domain.Transaction{}

					for list.Next(&transaction) {
						_, ok := storedDates[transaction.ID]

						So(ok, ShouldBeTrue)
						total++
					}

					Convey("And I can see the total of transactions inserted", func() {
						So(total, ShouldEqual, 45)
					})
				})
			})
			Convey("When I find the list based in a funding and a defined interval", func() {
				list := repository.FindByFundingAndInterval(transaction.Funding, initialDate.AddDate(0, 0, 30), initialDate.AddDate(0, 0, 35))
				total := 0

				Convey("Then I can iterate over all dates", func() {

					transaction := domain.Transaction{}
					lastDate := initialDate.AddDate(0, 0, 30)

					Convey("And the date is ordened", func() {
						for list.Next(&transaction) {
							So(transaction.Date, ShouldHappenOnOrAfter, lastDate)
							lastDate = transaction.Date
							total++
						}

						Convey("And I can see the total of transactions inserted", func() {
							So(total, ShouldEqual, 5)
						})
					})
				})
			})
		})
	})
}
