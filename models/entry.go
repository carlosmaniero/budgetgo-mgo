package models

import "time"

type Entry struct {
	Id      string
	Name    string
	Amount  float64
	Date    time.Time
	Comment string
}
