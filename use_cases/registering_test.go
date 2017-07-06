package use_cases

import (
	. "github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var repositoryCalls = 0

func TestSpec(t *testing.T) {
	Convey("Scenario: Registering a Transaction", t, func() {
		Convey("Given I've a Valid Transaction", func() {
			description := "4 beers"
			amount := 24.99

			Convey("When I register the transaction", func() {
				iterator := TransactionInteractor{Repository: &transactionRepository{}}
				err, _ := iterator.Register(description, amount)

				Convey("Then the register is saved successly", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repositoryCalls, ShouldEqual, 1)
				})
			})
		})
	})
}

// Mocked Repository
type transactionRepository struct{}

func (t *transactionRepository) Store(transaction *Transaction) {
	repositoryCalls = 1

	So(transaction.Description, ShouldEqual, "4 beers")
	So(transaction.Amount, ShouldEqual, 24.99)
}
