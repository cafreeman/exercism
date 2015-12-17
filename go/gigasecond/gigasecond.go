package gigasecond

import (
	"math"
	"time"
)

const (
	TestVersion = 2
)

var Birthday, _ = time.Parse("2006-01-02", "1986-08-18")
var Gigasecond = time.Duration(math.Pow(10, 9)) * time.Second

func AddGigasecond(in time.Time) time.Time {
	return in.Add(Gigasecond)
}
