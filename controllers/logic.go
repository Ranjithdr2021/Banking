package controllers

import t "time"

func isExpired(time transaction) bool {
	if duration := time.Timestamp.Time.Sub(t.Now()); duration > t.Duration(t.Duration(60).Seconds()) {
		return true
	}
	return false
}

//2021-07-17T09:59:51.312Z
