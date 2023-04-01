package controllers

import (
	t "time"
)

func isExpired(time []transaction) bool {
	// fmt.Println(t.Now().UTC())
	// fmt.Println(time.Timestamp.Time)
	// fmt.Println(60 * t.Second)
	// fmt.Println(t.Now().UTC().Sub(time.Timestamp.Time).Abs())
	return t.Now().UTC().Sub(time[0].Timestamp.Time) >= 60*t.Second
}

// 2021-07-17T09:59:51.312Z
func isInFuture(time transaction) bool {
	is := time.Timestamp.Time.After(t.Now())
	return is
}

//4:30:40 -- 4:30:10 =
