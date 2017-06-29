package handlers

import (
	. "github.com/carlosmaniero/budgetgo/controllers"
	. "github.com/carlosmaniero/budgetgo/models"
	. "github.com/carlosmaniero/budgetgo/serializers"
	"net/http"
)

func EntryCreateHandler(w http.ResponseWriter, r *http.Request) {
	entry := Entry{}
	serializer := EntrySerializer{&entry}
	serializer.Unserialize(r.Body)

	err := EntryCreateController(&entry)

	if err != nil {
		errorSerializer := EntryErrorSerializer{err}
		data, _ := errorSerializer.Serialize()
		w.Write(data)
		return
	}
	data, _ := serializer.Serialize()
	w.Write(data)
}
