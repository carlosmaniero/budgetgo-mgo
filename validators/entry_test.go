package validators

import (
	. "github.com/carlosmaniero/budgetgo/serializers"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func checkInvalidEntryField(t *testing.T, entry *EntryData, field string, errorMessage string) {
	Convey("Given I've an invalid "+field, t, func() {

		Convey("When I validate it", func() {
			validator := NewEntryValidator(entry)

			Convey("Then the validator returns that its is invalid", func() {
				So(validator.IsValid(), ShouldBeFalse)

				Convey("And show me that the amount should be a number", func() {
					errors := validator.GetErrors()
					So(errors[field], ShouldEqual, errorMessage)
				})
			})
		})
	})
}

func TestSpec(t *testing.T) {
	Convey("Given I've a valid EntryData", t, func() {
		entry := EntryData{
			Name:    "Bigode's Bakery",
			Amount:  "-10.30",
			Date:    "1993-09-25T05:30:01Z",
			Comment: "6 breads",
		}

		Convey("When I validate it", func() {
			validator := NewEntryValidator(&entry)

			Convey("Then the validator returns that its valid", func() {
				So(validator.IsValid(), ShouldBeTrue)
			})
		})
	})

	Convey("Given I've a blank EntryData", t, func() {
		entry := EntryData{}

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
						So(errors["Amount"], ShouldEqual, "This field is required.")
					})

					Convey("Then the validator show me that the field Date is required", func() {
						So(errors["Date"], ShouldEqual, "This field is required.")
					})
				})
			})

		})
	})

	checkInvalidEntryField(
		t,
		&EntryData{
			Name:    "Bigode's Bakery",
			Amount:  "it's no a number",
			Date:    "1993-09-25T05:30:01Z",
			Comment: "6 breads",
		},
		"Amount",
		"This field is not a number.",
	)

	checkInvalidEntryField(
		t,
		&EntryData{
			Name:    "Bigode's Bakery",
			Amount:  "10.5",
			Date:    "25/09/1993",
			Comment: "6 breads",
		},
		"Date",
		"This field is not a RFC3339 date.",
	)
}
