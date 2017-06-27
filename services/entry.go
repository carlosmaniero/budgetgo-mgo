package services

import (
	"errors"
	"fmt"
	"github.com/carlosmaniero/budgetgo/models"
)

type EntryService struct {
	entries []*models.Entry
}

func (service *EntryService) Insert(ientry interface{}) error {
	entry, converted := ientry.(*models.Entry)

	if !converted {
		return errors.New("There is not an Entry model at the first argument")
	}

	service.entries = append(service.entries, entry)
	_, total := service.Count()
	entry.Id = fmt.Sprint(total)
	return nil
}

func (service *EntryService) Count() (error, int) {
	return nil, len(service.entries)
}

func (service *EntryService) FindById(id string, ientry interface{}) error {
	entry2, converted := ientry.(*models.Entry)

	if !converted {
		return errors.New("There is not an Entry model at the first argument")
	}

	for _, entry := range service.entries {
		if entry.Id == id {
			entry2.Id = entry.Id
			entry2.Name = entry.Name
			entry2.Amount = entry.Amount
			entry2.Date = entry.Date
			entry2.Comment = entry.Comment

			return nil
		}
	}

	return errors.New("Element not found")
}

func NewEntryService() Service {
	return &EntryService{entries: []*models.Entry{}}
}
