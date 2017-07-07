package domain

import (
	. "github.com/smartystreets/goconvey/convey"
)

func shouldHaveErrorIn(errs []error, field string, message string) {
	founded := false
	var foundedError *FieldValidationError

	for _, err := range errs {
		validationError := err.(*FieldValidationError)
		if validationError.Field == field {
			founded = true
			foundedError = validationError
			break
		}
	}

	Convey("The "+field+" field has an error", func() {
		So(founded, ShouldBeTrue)
	})
	Convey("And the message error: "+message, func() {
		So(foundedError.Error(), ShouldEqual, message)
	})
}
