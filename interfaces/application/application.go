package application

import (
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/memoryrepository"
	"github.com/carlosmaniero/budgetgo/interfaces/repositories/mongorepository"
	"github.com/carlosmaniero/budgetgo/usecases"
	mgo "gopkg.in/mgo.v2"
)

// Application is the application context. This contains all repositories of
// the application.
type Application struct {
	TransactionRepository usecases.TransactionRepository
	FundingRepository     usecases.FundingRepository
}

// New create a new application instance
func New() *Application {
	return &Application{
		TransactionRepository: memoryrepository.NewMemoryTransactionRepository(),
		FundingRepository:     memoryrepository.NewMemoryFundingRepository(),
	}
}

// NewWithMongo returns a new application instance running with mongodb
func NewWithMongo(session *mgo.Session) *Application {
	db := session.DB("budgetgo")
	return &Application{
		TransactionRepository: mongorepository.NewMongoTransactionRepository(db.C("transaction")),
		FundingRepository:     mongorepository.NewMongoFundingRepository(db.C("funding")),
	}
}
