package controllers

import (
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/carlosmaniero/budgetgo/services"
	. "github.com/carlosmaniero/budgetgo/validators"
)

func EntryCreateController(entry *Entry) *EntryError {
	if entry.Id != "" {
		return &EntryError{Code: AlreadyCreatedError}
	}
	validator := NewEntryValidator(entry)
	ok := validator.IsValid()

	if !ok {
		return &EntryError{
			Code:        ValidationError,
			FieldErrors: validator.GetErrors(),
		}
	}

	service := NewEntryService()
	service.Insert(entry)

	return nil
}
