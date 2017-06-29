package handlers

import (
	. "github.com/carlosmaniero/budgetgo/controllers"
	"github.com/carlosmaniero/budgetgo/serializers"
	"net/http"
)

func EntryCreateHandler(w http.ResponseWriter, r *http.Request) {
	data, errS := serializers.StringToEntryData(r.Body)

	entry, err := EntryCreateController(&data)
	serializer := serializers.EntrySerializer{entry}

	if err != nil {
		errorSerializer := serializers.EntryErrorSerializer{err}
		data, _ := errorSerializer.Serialize()
		w.Write(data)
		return
	}
	responseBody, _ := serializer.Serialize()
	w.Write(responseBody)
}
