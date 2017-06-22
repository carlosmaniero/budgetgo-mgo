package models

import "time"

type Entry struct {
	Name    string
	amount  float32
	date    time.Date
	comment string
}
