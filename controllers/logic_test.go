package controllers

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/relvacode/iso8601"
)

func Test_IsExpired(t *testing.T) {
	tests := []struct {
		name string
		time transaction
		want bool
	}{
		{
			name: "Positive",
			time: transaction{Timestamp: iso8601.Time{time.Now()}},
			want: false,
		},
		{
			name: "Negetive",
			time: transaction{Timestamp: iso8601.Time{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isExpired(tt.time)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isInFuture(t *testing.T) {
	tests := []struct {
		name string
		time transaction
		want bool
	}{
		{
			name: "Positive",
			time: transaction{Timestamp: iso8601.Time{time.Now()}},
			want: false,
		},
		{
			name: "Negetive",
			time: transaction{Timestamp: iso8601.Time{time.Now().UTC().Add(60 * time.Second)}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInFuture(tt.time)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_statistic(t *testing.T) {
	tests := []struct {
		name     string
		transact []transaction
		want     *statistics
		want1    bool
	}{
		{
			name: "empty transactions",
			transact: []transaction{
				{
					Amount:    "100.43",
					Timestamp: iso8601.Time{Time: time.Now()},
				},
				{
					Amount:    "200.43",
					Timestamp: iso8601.Time{Time: time.Now()},
				},
				{
					Amount:    "50.43",
					Timestamp: iso8601.Time{Time: time.Now()},
				},
			},
			want:  &statistics{Sum: "351.29", Avg: "117.09666666666668", Max: "200.43", Min: "50.43", Count: "3"},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := statistic(tt.transact)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, isTrasact(tt.transact))
		})
	}
}
