package controllers

import (
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/serializers"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Given I've a valid Entry", t, func() {
		entry := EntryData{
			Name:    "Bigode's Bakery",
			Amount:  "-10.30",
			Date:    "1993-09-25T05:30:01Z",
			Comment: "6 breads",
		}

		Convey("When i send the entry to the EntryController", func() {
			entryModel, _ := EntryCreateController(&entry)

			Convey("Then the Entry was saved successly", func() {
				So(entryModel.Id, ShouldNotBeEmpty)
			})
		})
	})

	Convey("Given I've a Entry with ID", t, func() {
		entry := EntryData{
			Id:      "123",
			Name:    "Bigode's Bakery",
			Amount:  "-10.30",
			Date:    "1993-09-25T05:30:01Z",
			Comment: "6 breads",
		}

		Convey("When I send the entry to the EntryController", func() {
			_, err := EntryCreateController(&entry)

			Convey("Then the Controller raise an creation error", func() {
				So(err.Code, ShouldEqual, AlreadyCreatedError)
			})
		})
	})

	Convey("Given I've a invalid Entry", t, func() {
		entry := EntryData{}

		Convey("When I send the entry to the EntryController", func() {
			_, err := EntryCreateController(&entry)

			Convey("Then the Controller raise a validation error", func() {
				So(err.Code, ShouldEqual, ValidationError)
			})

			Convey("And need to have field errors", func() {
				So(len(err.FieldErrors), ShouldEqual, 3)
			})
		})
	})
}
