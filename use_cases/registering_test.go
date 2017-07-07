package use_cases

import (
	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Scenario: Registering a Transaction", t, func() {
		Convey("Given I've a Valid Transaction", func() {
			description := "4 beers"
			amount := 24.99

			Convey("When I register the transaction", func() {
				repository := transactionRepository{
					storedTotal:         0,
					expectedDescription: description,
					expectedAmount:      amount,
				}
				iterator := TransactionInteractor{Repository: &repository}
				err, _ := iterator.Register(description, amount)

				Convey("Then the transaction is saved successly", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 1)
				})
			})
		})

		Convey("Given I've a invalid Transaction", func() {
			description := ""
			amount := 0.0

			Convey("When I register the transaction", func() {
				repository := transactionRepository{
					storedTotal:         0,
					expectedDescription: description,
					expectedAmount:      amount,
				}
				iterator := TransactionInteractor{Repository: &repository}
				err, _ := iterator.Register(description, amount)

				Convey("Then the transaction isn't saved", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 0)
				})
			})
		})
	})
}

// Mocked Repository
type transactionRepository struct {
	storedTotal         int
	expectedDescription string
	expectedAmount      float64
}

func (t *transactionRepository) Store(transaction *domain.Transaction) {
	t.storedTotal++
	So(transaction.Description, ShouldEqual, t.expectedDescription)
	So(transaction.Amount, ShouldEqual, t.expectedAmount)
}
