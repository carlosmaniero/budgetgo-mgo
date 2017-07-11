package serializers

import (
	"encoding/json"
	"io"

	"github.com/carlosmaniero/budgetgo/domain"
)

type FundingData struct {
	Name       string  `json:"name"`
	Limit      float64 `json:"limit"`
	Amount     float64 `json:"amount"`
	ClosingDay int     `json:"closing_day"`
}

type FundingResponseData struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Limit      float64 `json:"limit"`
	Amount     float64 `json:"amount"`
	ClosingDay int     `json:"closing_day"`
}

type FundingValidationErrorData struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  []*FieldErrorData `json:"errors"`
}

func UnserializeFundingData(reader io.Reader) (*FundingData, error) {
	data := FundingData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}

func SerializeFunding(funding *domain.Funding) []byte {
	data := FundingResponseData{
		Id:         funding.Id,
		Name:       funding.Name,
		Amount:     funding.Amount,
		ClosingDay: funding.ClosingDay,
		Limit:      funding.Limit,
	}
	b, _ := json.Marshal(data)
	return b
}
func SerializeFundingValidationError(data *FundingValidationErrorData) []byte {
	b, _ := json.Marshal(data)
	return b
}
