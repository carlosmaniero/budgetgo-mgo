package serializers

import (
	"encoding/json"
	"io"

	"github.com/carlosmaniero/budgetgo/domain"
)

// FundingData is the serializable representation of a Funding Data
//
// This is used to parse the information sended to the Funding creation entrypoint
type FundingData struct {
	Name       string  `json:"name"`
	Limit      float64 `json:"limit"`
	Amount     float64 `json:"amount"`
	ClosingDay int     `json:"closing_day"`
}

// FundingResponseData is the serializable representation of a Funding
type FundingResponseData struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Limit      float64 `json:"limit"`
	Amount     float64 `json:"amount"`
	ClosingDay int     `json:"closing_day"`
}

// FundingValidationErrorData is the serializable representaion of a Funding
// validation error
type FundingValidationErrorData struct {
	Type    string            `json:"type"`
	Message string            `json:"message"`
	Errors  []*FieldErrorData `json:"errors"`
}

// UnserializeFundingData get an io.Reader and convert it to a FundingData
//
// This return an error if the input is not an valid json representation of
// the FundingData
func UnserializeFundingData(reader io.Reader) (*FundingData, error) {
	data := FundingData{}
	err := json.NewDecoder(reader).Decode(&data)
	return &data, err
}

// SerializeFunding receive a Funding and return its string representation
func SerializeFunding(funding *domain.Funding) []byte {
	data := FundingResponseData{
		ID:         funding.ID,
		Name:       funding.Name,
		Amount:     funding.Amount,
		ClosingDay: funding.ClosingDay,
		Limit:      funding.Limit,
	}
	b, _ := json.Marshal(data)
	return b
}

// SerializeFundingValidationError returns a json string representation of an
// FundingValidationErrorData
func SerializeFundingValidationError(data *FundingValidationErrorData) []byte {
	b, _ := json.Marshal(data)
	return b
}
