package isOdd

import (
	"strconv"
	"testing"
)

func TestIsOdd(t *testing.T) {
	for num := 0; num < 10000; num++ {
		res := IsOdd(num)
		number := strconv.Itoa(num)
		result := strconv.FormatBool(res)
		t.Log(number + " is odd: " + result)
	}
}
