package controllers

import "github.com/relvacode/iso8601"

type transaction struct {
	Amount    string       `json:"amount"`
	Timestamp iso8601.Time `json:"timestamp"`
}
