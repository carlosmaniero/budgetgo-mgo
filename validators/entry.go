package validators

import (
	. "github.com/carlosmaniero/budgetgo/models"
)

type EntryValidator struct {
	entry  *Entry
	errors map[string]string
}

func (validator *EntryValidator) IsValid() bool {
	validator.errors = make(map[string]string)
	validator.validateName()
	validator.validateAmount()

	return len(validator.errors) == 0
}

func (validator *EntryValidator) validateName() {
	if validator.entry.Name == "" {
		validator.errors["Name"] = "This field is required."
	}
}

func (validator *EntryValidator) validateAmount() {
	if validator.entry.Amount == 0 {
		validator.errors["Amount"] = "This field shouldn't be zero."
	}
}

func (validator *EntryValidator) GetErrors() map[string]string {
	return validator.errors
}

func NewEntryValidator(entry *Entry) *EntryValidator {
	return &EntryValidator{entry, make(map[string]string)}
}
