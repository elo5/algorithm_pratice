package easy

import (
	"fmt"
	"math"
)

func IsPalindrome(x int) bool {
	if x < 0  || x > math.MaxInt32 {
		return false
	}

	length := 0
	tempx := x

	for tempx >0 {
		tempx /=10
		length++
	}
	isP :=  true
	for i:= 0; i<length - 1; i++ {

		for j:= length -1; j>0; j-- {
			if i + j  == length - 1  { //&&  i<=j
				fmt.Println(i,j)
				//n/10^(site-1)%10
			}
		}
	}

	return isP
}