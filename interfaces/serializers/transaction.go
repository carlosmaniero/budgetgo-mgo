package serializers

import (
	"io"
	"encoding/json"
	"github.com/carlosmaniero/budgetgo/domain"
)


type TransactionData struct {
	Description string  `json:"description"`
	Amount float64 		`json:"amount"`
}


func UnserializeTransactionData(reader io.Reader) (*TransactionData, error) {
	data := TransactionData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}

func SerializeTransaction(transaction *domain.Transaction) []byte {
	data := TransactionData{
		Description: transaction.Description,
		Amount: transaction.Amount,
	}
	b, _ := json.Marshal(data)
	return b
}