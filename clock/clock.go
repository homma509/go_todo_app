package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClock struct{}

func (r RealClock) Now() time.Time {
	return time.Now()
}

type FixedClock struct{}

func (fc FixedClock) Now() time.Time {
	return time.Date(2025, 1, 6, 7, 3, 56, 0, time.UTC)
}
