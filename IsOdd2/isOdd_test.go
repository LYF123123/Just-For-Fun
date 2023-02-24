package isodd2

import "testing"

func TestIsEven(t *testing.T) {
	a := 100
	t.Log(IsEven(a))
	t.Log(IsOdd(a))
}
