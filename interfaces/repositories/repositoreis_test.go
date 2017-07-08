package repositories

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/carlosmaniero/budgetgo/domain"
	"github.com/carlosmaniero/budgetgo/usecases"
)

func TestSpecFounding(t *testing.T) {
	Convey("Scenario: Registering an TransactionRepository", t, func() {
		Convey("Given I've Dummy TransactionRepository", func() {
			Convey("When I register it", func() {
				RegisterTransactionRepository("dummy", newTransactionRepository)

				Convey("Then I can get the registered repository", func() {
					repository, err := NewTransactionRepository("dummy")
					So(err, ShouldBeNil)

					Convey("And it's a dummy repository", func() {
						_, ok :=repository.(*transactionRepository)

						So(ok, ShouldBeTrue)
					})
				})
			})
		})
	})
}

type transactionRepository struct {}
func (t *transactionRepository) Store(transaction *domain.Transaction) {}

func newTransactionRepository() usecases.TransactionRepository {
	return &transactionRepository{}
}