package easy

import (
	"strconv"
	"strings"
)

/**
415. Add Strings
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
 */

func AddStrings(num1 string, num2 string) string {

	len1 := len(num1) - 1
	len2 := len(num2) - 1

	passThrough := 0
	index := 0

	res := strings.Builder{}

	for len1 >= 0 || len2 >= 0 {

		temp := passThrough

		if len1 >=0 {
			temp += int(num1[len1] - '0')
		}

		if len2 >= 0 {
			temp += int(num2[len2] - '0')
		}

		if temp > 9 {
			 res.WriteString(strconv.Itoa(temp - 10))
			 passThrough = 1
		}else {
			res.WriteString(strconv.Itoa(temp ))
			passThrough = 0
		}

		len1--
		len2--
		index++

	}

	if passThrough != 0 {
		res.WriteString("1")
	}

	return reverseStr(res.String())
}

func reverseStr(s string) string  {

	r := []rune(s)

	for i,j := 0, len(r) -1; i< j; i,j = i+1, j-1 {
		r[i], r[j] = r[j],r[i]
	}
	return string(r)
}
