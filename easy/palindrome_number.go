package easy

import (
	"math"
)

//IsPalindrome 数字是否是对称的
//排除数字超过范围后，reverse integer后的值是否跟原来的数字相等
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