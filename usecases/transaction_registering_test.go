package usecases

import (
	"strconv"
	"testing"
	"time"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Scenario: Registering a Transaction", t, func() {
		Convey("Given I've a Valid Transaction", func() {
			description := "4 beers"
			amount := 24.99
			date := time.Now()
			funding := domain.Funding{
				Name:       "Bank account",
				Limit:      1000,
				Amount:     0,
				ClosingDay: 1,
			}

			Convey("When I register the transaction", func() {
				repository := transactionRepository{
					storedTotal:         0,
					expectedDescription: description,
					expectedAmount:      amount,
					expectedDate:        date,
					expectedFunding:     funding,
				}
				iterator := TransactionInteractor{Repository: &repository}
				err, transaction := iterator.Register(description, amount, date, funding)

				Convey("Then the transaction is saved successly", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 1)
				})

				Convey("And the transaction has the created id", func() {
					So(transaction.Id, ShouldEqual, "1")
				})
			})
		})

		Convey("Given I've a invalid Transaction", func() {
			description := ""
			amount := 0.0
			date := time.Now().AddDate(0, -1, -1)
			funding := domain.Funding{}

			Convey("When I register the transaction", func() {
				repository := transactionRepository{
					storedTotal:         0,
					expectedDescription: description,
					expectedDate:        date,
					expectedAmount:      amount,
				}
				iterator := TransactionInteractor{Repository: &repository}
				err, _ := iterator.Register(description, amount, date, funding)

				Convey("Then the transaction isn't saved", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("And the data isn't saved inside the repository", func() {
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
	expectedDate        time.Time
	expectedFunding     domain.Funding
}

func (t *transactionRepository) Store(transaction *domain.Transaction) string {
	t.storedTotal++
	So(transaction.Description, ShouldEqual, t.expectedDescription)
	So(transaction.Amount, ShouldEqual, t.expectedAmount)
	return strconv.Itoa(t.storedTotal)
}
