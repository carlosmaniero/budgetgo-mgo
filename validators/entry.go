package validators

import (
	. "github.com/carlosmaniero/budgetgo/serializers"
	"strconv"
	"time"
)

type EntryValidator struct {
	entry  *EntryData
	errors map[string]string
}

func (validator *EntryValidator) IsValid() bool {
	validator.errors = make(map[string]string)

	validator.validateName()
	validator.validateAmount()
	validator.validateDate()

	return len(validator.errors) == 0
}

func (validator *EntryValidator) checkRequired(fieldName string, value string) bool {
	if value == "" {
		validator.errors[fieldName] = "This field is required."
		return false
	}

	return true
}

func (validator *EntryValidator) validateName() {
	validator.checkRequired("Name", validator.entry.Name)
}

func (validator *EntryValidator) validateAmount() {
	ok := validator.checkRequired("Amount", validator.entry.Amount)

	if !ok {
		return
	}

	_, err := strconv.ParseFloat(validator.entry.Amount, 32)

	if err != nil {
		validator.errors["Amount"] = "This field is not a number."
	}
}

func (validator *EntryValidator) validateDate() {
	ok := validator.checkRequired("Date", validator.entry.Date)

	if !ok {
		return
	}

	_, err := time.Parse(time.RFC3339, validator.entry.Date)

	if err != nil {
		validator.errors["Date"] = "This field is not a RFC3339 date."
	}
}

func (validator *EntryValidator) GetErrors() map[string]string {
	return validator.errors
}

func NewEntryValidator(entry *EntryData) *EntryValidator {
	return &EntryValidator{entry, make(map[string]string)}
}
