package isOdd

import (
	"math"
)

func IsOdd(n int) bool {
	n = int(math.Abs(float64(n)))
	if n == 0 {
		return false
	}
	return !IsOdd(n - 1)
}
