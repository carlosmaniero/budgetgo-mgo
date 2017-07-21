package mongorepository

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/carlosmaniero/budgetgo/domain"
	. "github.com/smartystreets/goconvey/convey"
)

var funding = domain.Funding{
	ID:         bson.NewObjectId().Hex(),
	Amount:     100,
	ClosingDay: 1,
	Limit:      10,
	Name:       "Bank account",
}

func TestCaseFunding(t *testing.T) {
	Convey("Scenario: Registering a funding", t, func() {
		repository := NewMongoFundingRepository(db.C("test-funding-on-create"))
		Convey("Given I've a funding", func() {
			Convey("When I insert it", func() {
				id := repository.Store(&funding)

				Convey("Then I receive an ObjectId as string", func() {
					So(bson.IsObjectIdHex(id), ShouldBeTrue)
				})
			})
		})
	})
	Convey("Scenario: Get a created funding", t, func() {
		repository := NewMongoFundingRepository(db.C("test-funding-on-retrieve"))

		Convey("Given I've a created funding", func() {
			id := repository.Store(&funding)

			Convey("When I get it", func() {
				fundingReceived := repository.FindByID(id)

				Convey("Then I've the same funding", func() {
					So(fundingReceived.ID, ShouldEqual, id)
					So(fundingReceived.Amount, ShouldEqual, funding.Amount)
					So(fundingReceived.Name, ShouldEqual, funding.Name)
					So(fundingReceived.ClosingDay, ShouldEqual, funding.ClosingDay)
				})
			})
		})
	})
}
