package leap

import "math"

func IsLeapYear(year int) bool {
	var yearFloat float64 = float64(year)
	if math.Mod(yearFloat, 400) == 0 {
		return true
	}
	if math.Mod(yearFloat, 100) == 0 {
		return false
	}
	if math.Mod(yearFloat, 4) == 0 {
		return true
	}
	return false
}

const TestVersion = 1
