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
			transaction := domain.Transaction{
				Description: "4 beers",
				Amount:      24.99,
				Date:        time.Now(),
				Funding: &domain.Funding{
					ID:         "funding-id",
					Name:       "Bank account",
					Limit:      1000,
					Amount:     0,
					ClosingDay: 1,
				},
			}

			Convey("When I register the transaction", func() {
				repository := transactionRepository{storedTotal: 0, expectedTransaction: &transaction}
				iterator := TransactionInteractor{Repository: &repository}
				err := iterator.Register(&transaction)

				Convey("Then the transaction is saved successly", func() {
					So(err, ShouldBeNil)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 1)
				})

				Convey("And the transaction has the created id", func() {
					So(transaction.ID, ShouldEqual, "1")
				})
			})
		})

		Convey("Given I've a invalid Transaction", func() {
			transaction := domain.Transaction{
				Description: "",
				Amount:      0.0,
				Date:        time.Now().AddDate(0, -1, -1),
				Funding:     &domain.Funding{},
			}

			Convey("When I register the transaction", func() {
				repository := transactionRepository{storedTotal: 0, expectedTransaction: &transaction}
				iterator := TransactionInteractor{Repository: &repository}
				err := iterator.Register(&transaction)

				Convey("Then the transaction isn't saved", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("And the data isn't saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 0)
				})
			})
		})
	})
	Convey("Scenario: Consulting a transaction", t, func() {
		Convey("Given I've a created transaction", func() {
			createdTransaction := domain.Transaction{
				Description: "4 beers",
				Amount:      24.99,
				Date:        time.Now(),
				Funding: &domain.Funding{
					ID:         "funding-id",
					Name:       "Bank account",
					Limit:      1000,
					Amount:     0,
					ClosingDay: 1,
				},
			}
			repository := transactionRepository{storedTotal: 0, expectedTransaction: &createdTransaction}
			iterator := TransactionInteractor{Repository: &repository}
			iterator.Register(&createdTransaction)

			Convey("When I retrieve it", func() {
				transaction, _ := iterator.Retrieve(createdTransaction.ID)

				Convey("Then I receive my created transaction", func() {
					So(transaction, ShouldEqual, &createdTransaction)
				})
			})

			Convey("When I retrieve a nonexistent transaction", func() {
				repository.lastStored = nil
				transaction, err := iterator.Retrieve("invalid-id")

				Convey("Then I receive my created transaction", func() {
					So(transaction, ShouldBeNil)
					So(err, ShouldEqual, ErrTransactionNotFound)
				})
			})
		})
	})
	Convey("Scenario: Consulting the list of transactions in a month", t, func() {
		Convey("Given I've a valid date and funding", func() {
			Convey("When I try to get the transaction list", func() {
				repository := transactionRepository{}
				iterator := TransactionInteractor{Repository: &repository}
				funding := domain.Funding{
					ID:         "funding-id",
					Name:       "Bank account",
					Limit:      1000,
					Amount:     0,
					ClosingDay: 1,
				}
				iterator.RetriveByFundingMonth(&funding, 2017, 6)

				Convey("Then the use case query correctly the repository", func() {
					So(repository.findedFunding, ShouldEqual, &funding)
					So(repository.findedStart.Unix(), ShouldEqual, time.Date(2017, time.Month(6), 1, 0, 0, 0, 0, time.Local).Unix())
					So(repository.findedEnd.Unix(), ShouldEqual, time.Date(2017, time.Month(6), 30, 0, 0, 0, 0, time.Local).Unix())
				})
			})
		})

		Convey("Given I've a invalid month", func() {
			Convey("When I try to get the transaction list", func() {
				repository := transactionRepository{}
				iterator := TransactionInteractor{Repository: &repository}
				funding := domain.Funding{
					ID:         "funding-id",
					Name:       "Bank account",
					Limit:      1000,
					Amount:     0,
					ClosingDay: 1,
				}
				_, err1 := iterator.RetriveByFundingMonth(&funding, 2017, 13)
				_, err2 := iterator.RetriveByFundingMonth(&funding, 2017, 0)

				Convey("When an error is returned", func() {
					So(err1, ShouldEqual, ErrInvalidMonth)
					So(err2, ShouldEqual, ErrInvalidMonth)
				})
			})
		})
	})
}

// Mocked Repository
type transactionRepository struct {
	storedTotal         int
	expectedTransaction *domain.Transaction
	lastStored          *domain.Transaction
	findedFunding       *domain.Funding
	findedStart         time.Time
	findedEnd           time.Time
}

func (t *transactionRepository) Store(transaction *domain.Transaction) string {
	t.storedTotal++
	So(transaction, ShouldEqual, t.expectedTransaction)
	t.lastStored = transaction
	return strconv.Itoa(t.storedTotal)
}

func (t *transactionRepository) FindByID(string) *domain.Transaction {
	return t.lastStored
}

func (t *transactionRepository) FindByFundingAndInterval(funding *domain.Funding, start time.Time, end time.Time) TransactionList {
	t.findedFunding = funding
	t.findedStart = start
	t.findedEnd = end
	return nil
}
