package controllers

import (
	. "github.com/carlosmaniero/budgetgo/errors"
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/carlosmaniero/budgetgo/serializers"
	. "github.com/carlosmaniero/budgetgo/services"
	. "github.com/carlosmaniero/budgetgo/validators"
)

func EntryCreateController(entryData *EntryData) (*Entry, *EntryError) {
	if entryData.Id != "" {
		return nil, &EntryError{Code: AlreadyCreatedError}
	}
	validator := NewEntryValidator(entryData)
	ok := validator.IsValid()

	if !ok {
		return nil, &EntryError{
			Code:        ValidationError,
			FieldErrors: validator.GetErrors(),
		}
	}

	entry := EntryDataToEntry(entryData)
	service := NewEntryService()
	service.Insert(entry)

	return entry, nil
}
