package basics

import (
	"fmt"
	"math"
)

func Sqrt() string {
	x := -4.0
	if x < 0 {
		return "NaN"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func PowWithLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
