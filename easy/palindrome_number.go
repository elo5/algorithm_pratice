package easy

import (
	"math"
)

func IsPalindrome(x int) bool {
	if x < 0  || x > math.MaxInt32 {
		return false
	}
	preserveX := x
	reversedX := 0
	for x > 0 {
		remain := x % 10
		reversedX *= 10
		reversedX += remain
		x /= 10
	}
	if preserveX == reversedX {
		return true
	}
	return false
}