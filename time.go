package goja

import (
	"time"
)

var (
	myTime = &Time{
		Second:      time.Second,
		Hour:        time.Hour,
		Minute:      time.Minute,
		Microsecond: time.Microsecond,
		Millisecond: time.Millisecond,
	}
)

type Time struct {
	Second      time.Duration
	Hour        time.Duration
	Minute      time.Duration
	Microsecond time.Duration
	Millisecond time.Duration
}

func (p Time) Now() time.Time {
	return time.Now()
}

func (p Time) Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}
