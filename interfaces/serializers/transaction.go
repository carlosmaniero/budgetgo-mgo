package serializers

import (
	"io"
	"encoding/json"
)


type TransactionData struct {
	Description string
	Amount float64
}


func UnserializeTransactionData(reader io.Reader) (*TransactionData, error) {
	data := TransactionData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}