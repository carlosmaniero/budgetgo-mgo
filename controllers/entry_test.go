package controllers

import (
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSpec(t *testing.T) {
	Convey("Given I've a valid Entry", t, func() {
		entry := Entry{
			Name:    "Bigode's Bakery",
			Amount:  -10.30,
			Date:    time.Now(),
			Comment: "6 breads",
		}

		Convey("When i send the entry to the EntryController", func() {
			EntryCreateController(&entry)

			Convey("Then the Entry was saved successly", func() {
				So(entry.Id, ShouldNotBeEmpty)
			})
		})
	})

	Convey("Given I've a Entry with ID", t, func() {
		entry := Entry{
			Id:      "123",
			Name:    "Bigode's Bakery",
			Amount:  -10.30,
			Date:    time.Now(),
			Comment: "6 breads",
		}

		Convey("When I send the entry to the EntryController", func() {
			err := EntryCreateController(&entry)

			Convey("Then the Controller raise an creation error", func() {
				So(err.Code, ShouldEqual, AlreadyCreatedError)
			})
		})
	})

	Convey("Given I've a invalid Entry", t, func() {
		entry := Entry{}

		Convey("When I send the entry to the EntryController", func() {
			err := EntryCreateController(&entry)

			Convey("Then the Controller raise a validation error", func() {
				So(err.Code, ShouldEqual, ValidationError)
			})

			Convey("And need to have field errors", func() {
				So(len(err.FieldErrors), ShouldEqual, 2)
			})
		})
	})
}
