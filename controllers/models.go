package controllers

import (
	"github.com/relvacode/iso8601"
)

type transaction struct {
	Amount    string       `json:"amount"`
	Timestamp iso8601.Time `json:"timestamp"`
}
type location struct {
	City string `json:"city"`
}

type reset struct {
	Reset bool
}
type statistics struct {
	Sum   string `json:"sum"`
	Avg   string `json:"avg"`
	Max   string `json:"max"`
	Min   string `json:"min"`
	Count string `json:"count"`
}
