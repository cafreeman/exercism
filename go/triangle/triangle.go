package triangle

import "math"

type Kind string

const (
	Sca = "Scalene"
	Iso = "Isosceles"
	Equ = "Equilateral"
	NaT = "Not a Triangle"
)

func KindFromSides(a, b, c float64) Kind {
	if cannotExist(a, b, c) || triangleInequality(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}

func cannotExist(a, b, c float64) bool {
	// Check each individual side to see if it's non-existent
	res := false
	for _, elem := range []float64{a, b, c} {
		if math.IsNaN(elem) || math.IsInf(elem, 1) || math.IsInf(elem, -1) {
			res = true
		}
	}
	return res
}

func triangleInequality(a, b, c float64) bool {
	res := false
	if a+b <= c || a+c <= b || b+c <= a {
		res = true
	}
	return res
}
