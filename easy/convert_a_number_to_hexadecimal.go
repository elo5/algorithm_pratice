package easy

import (
	"fmt"
	"strconv"
)

/*
405. Convert a Number to Hexadecimal
Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:
All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.

Example 1:
Input:
26
Output:
"1a"

Example 2:
Input:
-1
Output:
"ffffffff"
 */
func ToHex1(num int) string {
	if num == 0 {
		return "0"
	}

	arr := []byte {'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f'}

	if num < 0{
		num = -num
		num ^= 0xffffffff
		num += 1
	}

	res := ""

	for num > 0 {
		digit := num % 16
		res = string(arr[digit]) + res
		num /= 16
	}

	return res
}



func ToHex(num int) string {


	//if num == 0 {
	//	return "0"
	//}
	//
	//var i int
	//bytes := []byte("0123456789abcdef        ") // 8 spaces
	//
	//for i = 23; num != 0; i-- {
	//	bytes[i] = bytes[num&15]
	//	num = int(uint32(num) >> 4)
	//}
	//
	//return string(bytes[i+1:])


	if num == 0 {
		 return "0"
	}

	result := ""

	if num < 0{
		num = -num
		num ^= 0xffffffff
		num += 1
	}

	n := num

	numsMap := make(map[int]string)
	numsMap[10] = "a"
	numsMap[11] = "b"
	numsMap[12] = "c"
	numsMap[13] = "d"
	numsMap[14] = "e"
	numsMap[15] = "f"

	for n >= 16 {
		m := n % 16
		fmt.Println("\n######")
		//fmt.Println("整除结果： " , n, "   取余", m)
		n = n / 16
		fmt.Println("整除结果： " , n, "   取余", m)
		fmt.Println("整除结果： " , n)
		if m > 9 {
			result = numsMap[m] + result
		}else {
			result = strconv.Itoa(m) + result
		}
	}
	if n != 0 {
		if n > 9 {
			result = numsMap[n] + result
		}else {
			result = strconv.Itoa(n) + result
		}
	}
	return result
}
