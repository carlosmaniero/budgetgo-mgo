package validators

import (
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

		Convey("When I validate it", func() {
			validator := NewEntryValidator(&entry)

			Convey("Then the validator returns that its valid", func() {
				So(validator.IsValid(), ShouldBeTrue)
			})
		})
	})

	Convey("Given I've an invalid Entry", t, func() {
		entry := Entry{Amount: 0.0}

		Convey("When I validate it", func() {
			validator := NewEntryValidator(&entry)

			Convey("Then the validator returns that its is invalid", func() {
				So(validator.IsValid(), ShouldBeFalse)

				Convey("And when I see the error list", func() {
					errors := validator.GetErrors()

					Convey("Then the validator show me that the field name is required", func() {
						So(errors["Name"], ShouldEqual, "This field is required.")
					})

					Convey("Then the validator show me that the field amount is required", func() {
						So(errors["Amount"], ShouldEqual, "This field shouldn't be zero.")
					})
				})
			})

		})
	})
}
