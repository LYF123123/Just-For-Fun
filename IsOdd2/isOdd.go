package isodd2

func IsEven(num int) bool {
	if num == 0 {
		return true
	} else if num > 0 {
		return IsOdd(num - 1)
	} else {
		return IsOdd(num + 1)
	}
}

func IsOdd(num int) bool {
	if num == 0 {
		return false
	} else if num > 0 {
		return IsEven(num - 1)
	} else {
		return IsEven(num + 1)
	}
}
