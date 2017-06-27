package main

import "github.com/carlosmaniero/budgetgo/services"
import "github.com/carlosmaniero/budgetgo/models"
import "fmt"

func main() {
	service := services.NewEntryService()
	service.Insert(&models.Entry{})
	_, total := service.Count()
	fmt.Println(total)
}
