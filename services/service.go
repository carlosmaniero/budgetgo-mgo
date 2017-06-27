package services

type Service interface {
	Insert(interface{}) error
	FindById(string, interface{}) error
	Count() (error, int)
}
