package util

import "time"

type Time interface {
	Now() time.Time
}

type TimeWrapper struct {
}

func (w *TimeWrapper) Now() time.Time {
	return time.Now()
}
