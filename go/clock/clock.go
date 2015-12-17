package clock

import "fmt"

const (
	minutesPerDay  = 1440
	minutesPerHour = 60
	TestVersion    = 2
)

//Clock export
type Clock int

//Time constructor for Clock
func Time(hour, minute int) Clock {
	return Clock(hour * minutesPerHour).Add(minute)
}

func (clock Clock) String() string {
	hours := clock / minutesPerHour
	minutes := clock % minutesPerHour
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

//Add minutes to an existing clock
func (clock Clock) Add(minutes int) Clock {
	clock += Clock(minutes)
	return standardize(clock)
}

func standardize(clock Clock) Clock {
	clock %= minutesPerDay
	if clock < 0 {
		clock += minutesPerDay
	}
	return clock
}
