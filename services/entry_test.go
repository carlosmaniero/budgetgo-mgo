package services

import (
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSpec(t *testing.T) {

	Convey("Given I've a debit entry", t, func() {
		entry := Entry{
			Name:    "Bigode's Bakery",
			Amount:  -10.30,
			Date:    time.Now(),
			Comment: "6 breads",
		}

		Convey("When I sent it to a service", func() {
			service := NewEntryService()
			err := service.Insert(&entry)

			Convey("There is no error given", func() {
				So(err, ShouldBeNil)
			})

			Convey("And the service creates an id inside the Entry", func() {
				So(entry.Id, ShouldNotBeNil)
				So(entry.Id, ShouldNotBeEmpty)
			})

			Convey("And When I find the created entry using the service", func() {
				entry2 := Entry{}
				service.FindById(entry.Id, &entry2)

				Convey("Then I've the same entry inserted", func() {
					So(entry.Id, ShouldEqual, entry2.Id)
					So(entry.Name, ShouldEqual, entry2.Name)
					So(entry.Amount, ShouldEqual, entry2.Amount)
					So(entry.Date, ShouldHappenOnOrBetween, entry2.Date, entry2.Date)
					So(entry.Comment, ShouldEqual, entry2.Comment)
				})
			})
		})
	})

	Convey("Given I've an invalid entry id", t, func() {
		id := "invalid-id"

		Convey("When I find it", func() {
			service := NewEntryService()
			err := service.FindById(id, &Entry{})

			Convey("There return an error message", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
