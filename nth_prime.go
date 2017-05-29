package prime

import (
	"math"
)

const testVersion = 1

func Nth(nth int) (int, bool) {
	if nth == 1 {
		return 2, true
	} else if nth == 0 {
		return 0, false
	}
	n := float64(float64(nth)*math.Log(float64(nth)) + float64(nth)*math.Log(math.Log(float64(nth))) + 6)
	a := make([]bool, int(n))

	for i := 2; i < int(math.Sqrt(n)); i++ {
		if !a[int(i)] {
			for j := i * i; j < int(n); j = j + i {
				a[j] = true
			}
		}
	}
	var p int
	for i := 2; i < int(n) && nth > 0; i++ {
		if !a[i] {
			p = i
			nth--
		}
	}
	return p, true
}
