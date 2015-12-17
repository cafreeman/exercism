package raindrops

import "strconv"

func Convert(num int) string {
	s := ""
	if num%3 == 0 {
		s += "Pling"
	}
	if num%5 == 0 {
		s += "Plang"
	}
	if num%7 == 0 {
		s += "Plong"
	}
	if len(s) == 0 {
		s += strconv.Itoa(num)
	}
	return s
}
