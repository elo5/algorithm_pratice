package main

import (
	"algorithm_pratice/easy"
	"fmt"
)

func main()  {

	//medium.Calculate("1 + 2 * 3")
	//fmt.Println("abc")

	//fmt.Println(easy.ReverseInteger_COOLFUN(-1234567))
	//fmt.Println(easy.IsPalindrome(123321))
	//fmt.Println(easy.TwoSum([]int{1,7,2,15},9))
	//fmt.Println(easy.TwoSum_faster([]int{3,5,4,2},6))
	fmt.Println(easy.TwoSum_faster([]int{2,4,7,9,8},12))
	fmt.Println(easy.TwoSum([]int{2,4,7,9,8},12))
}
