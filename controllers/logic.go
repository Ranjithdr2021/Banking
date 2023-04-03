package controllers

import (
	"fmt"
	"strconv"
	t "time"
)

func isExpired(time transaction) bool {
	// fmt.Println(t.Now().UTC())
	// fmt.Println(time.Timestamp.Time)
	// fmt.Println(60 * t.Second)
	// fmt.Println(t.Now().UTC().Sub(time.Timestamp.Time).Abs())
	duration := t.Now().UTC().Sub(time.Timestamp.Time)
	return duration >= 60*t.Second
}

// 2021-07-17T09:59:51.312Z
func isInFuture(time transaction) bool {
	is := time.Timestamp.Time.After(t.Now())
	return is
}

func statistic(transact []transaction) *statistics {
	var sum float64
	min, _ := strconv.ParseFloat(transact[0].Amount, 64)
	max := min
	for _, v := range transact { // 2 9 5 8 1
		v1, _ := strconv.ParseFloat(v.Amount, 64)
		if v1 > max {
			max = v1
		} else if v1 < min {
			min = v1
		}
		s, _ := strconv.ParseFloat(v.Amount, 64)
		sum += s
	}
	avg := (sum / float64(len(transact)))
	statistic := statistics{
		Sum:   fmt.Sprintf("%v", sum),
		Avg:   fmt.Sprintf("%v", avg),
		Max:   fmt.Sprintf("%v", max),
		Min:   fmt.Sprintf("%v", min),
		Count: fmt.Sprintf("%v", len(transact)),
	}
	return &statistic
}

func isTrasact(t []transaction) bool {
	return len(t) != 0
}
