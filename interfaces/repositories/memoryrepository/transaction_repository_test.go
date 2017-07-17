package memoryrepository

import (
	"testing"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecTransactionRepository(t *testing.T) {
	Convey("Scenario: Insert register asynchronously", t, func() {
		Convey("Given I've a list of transactions", func() {
			total := 0
			repo := MemoryTransactionRepository{transactions: make([]*domain.Transaction, 0)}
			concurrency := 10000
			sent := make(chan bool, 500)

			for i := 0; i < concurrency; i++ {
				go func() {
					repo.Store(&domain.Transaction{})
					sent <- true
				}()
			}

			for {
				<-sent
				total++
				if total == concurrency {
					So(len(repo.transactions), ShouldEqual, concurrency)
					return
				}
			}
		})
	})
}
