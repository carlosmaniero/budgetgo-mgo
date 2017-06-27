package models

import "time"

type Entry struct {
	Id      string
	Name    string
	Amount  float32
	Date    time.Time
	Comment string
}
