package clock

import "time"

//go:generate mockery --name Clock
type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}
