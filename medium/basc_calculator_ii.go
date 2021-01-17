package medium

import (
	"fmt"
	"strings"
)

/*
https://leetcode.com/problems/basic-calculator-ii/
227. Basic Calculator II   Medium
Given a string s which represents an expression, evaluate this expression and return its value.
The integer division should truncate toward zero.

Example 1:
Input: s = "3+2*2"
Output: 7

Example 2:
Input: s = " 3/2 "
Output: 1

Example 3:
Input: s = " 3+5 / 2 "
Output: 5

Constraints:
1 <= s.length <= 3 * 105
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 231 - 1].
The answer is guaranteed to fit in a 32-bit integer.
*/

func Calculate(s string) int {

	if len(s) == 0 ||  len(s) > 300000 {
		return 0
	}

	fmt.Println(s)
	//"1 - 2 + 2 * 3 /3 * 4"

	// 1 -2  + 2 ,  3 /3 , 4

	// 1 -2  + 2, 3, 3, 3

	//var nums []int
	//var op  = "+"
	//var cur = 0
	//var count = 0
	//for {
	//	if count >= len(s) {
	//		break
	//	}
	//
	//	if s[count] ==  ' '{
	//		continue
	//	}
	//
	//
	//
	//
	//
	//	fmt.Println("hello,world", count)
	//	count++
	//}




	splitedMulti := strings.Split(s,"*")

	if len(splitedMulti) > 0 {
		for i, vM  := range splitedMulti{
			fmt.Println(i,"  --  ",vM)
			splitedDivider :=  strings.Split(vM, "/")
			for j, vD := range splitedDivider{
				fmt.Println(j,"  --  ",vD)



			}

		}

	}


	result :=  0


	return result
}


