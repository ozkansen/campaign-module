package time

import (
	"time"
)

type (
	Duration = time.Duration
	Time     = time.Time
)

var (
	Hour = time.Hour
	Now  = time.Now
)

func Set(t time.Time) {
	Now = func() time.Time {
		return t
	}
}

func Increase(h int) {
	Set(Now().Add(Duration(h) * time.Hour))
}
