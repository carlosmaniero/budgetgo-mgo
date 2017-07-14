package usecases

import (
	"strconv"
	"testing"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecFundingRegistering(t *testing.T) {
	Convey("Scenario: Creating a funding", t, func() {
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
					So(funding.ID, ShouldEqual, "1")
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

func TestSpecFundingRetrieve(t *testing.T) {
	Convey("Scenario: Retrieve an funding", t, func() {
		repository := fundingRepository{storedTotal: 0}
		iteractor := FundingInteractor{&repository}
		Convey("Given I've a registred funding", func() {
			fundingCreated, _ := iteractor.Register("Beer account", 10.0, 10, 10.0)

			Convey("When I retrieve the registred transaction", func() {
				fundingRetrieved, _ := iteractor.Retrieve(fundingCreated.ID)

				Convey("Then it is returned", func() {
					So(fundingRetrieved, ShouldEqual, fundingCreated)
					So(fundingCreated.ID, ShouldEqual, repository.findedID)
				})
			})
		})
		Convey("Given I've a unregistred funding", func() {
			repository.stored = nil

			Convey("When I try to retrieve it", func() {
				fundingRetrieved, err := iteractor.Retrieve("id-not-found")

				Convey("Then need to return an error", func() {
					So(fundingRetrieved, ShouldBeNil)
					So(err, ShouldEqual, ErrFundingNotFound)
				})
			})
		})
	})
}

// Mocked Repository
type fundingRepository struct {
	storedTotal int
	stored      *domain.Funding
	findedID    string
}

func (f *fundingRepository) FindByID(id string) *domain.Funding {
	f.findedID = id
	return f.stored
}

func (f *fundingRepository) Store(funding *domain.Funding) string {
	f.storedTotal++
	f.stored = funding
	return strconv.Itoa(f.storedTotal)
}
