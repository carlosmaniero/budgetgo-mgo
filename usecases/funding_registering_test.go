package usecases

import (
	"strconv"
	"testing"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecFundingRegistering(t *testing.T) {
	Convey("Scenario: Creating an funding", t, func() {
		Convey("Given I've a valid funding", func() {
			name := "Bank account"
			amount := 10.0
			closingDay := 1
			limit := 1000.0

			Convey("When I register it", func() {
				repository := fundingRepository{storedTotal: 0}
				iterator := FundingInteractor{&repository}
				funding, err := iterator.Register(name, amount, closingDay, limit)

				Convey("Then The funding was registered succesfully", func() {
					So(err, ShouldBeNil)
					So(funding.Name, ShouldEqual, name)
					So(funding.Amount, ShouldEqual, amount)
					So(funding.ClosingDay, ShouldEqual, closingDay)
					So(funding.Limit, ShouldEqual, limit)
				})

				Convey("And the data is saved inside the repository", func() {
					So(repository.storedTotal, ShouldEqual, 1)
				})

				Convey("And the funding has the created id", func() {
					So(funding.Id, ShouldEqual, "1")
				})
			})
		})

		Convey("Given I've a invalid funding", func() {
			name := ""
			amount := 10.0
			closingDay := -1
			limit := 1000.0

			Convey("When I register it", func() {
				repository := fundingRepository{storedTotal: 0}
				iterator := FundingInteractor{&repository}
				_, err := iterator.Register(name, amount, closingDay, limit)

				Convey("Then The funding was registered succesfully", func() {
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
type fundingRepository struct {
	storedTotal int
}

func (t *fundingRepository) Store(funding *domain.Funding) string {
	t.storedTotal++
	return strconv.Itoa(t.storedTotal)
}
