package easy

import (
	"fmt"
	"math"
	"strconv"
)

//ReverseInteger_COOLFUN 超过范围返回0
//若是负数，保存一个布尔变量，返回结果前作转换
//除10取余的值*10 （表达能力有限，代码更直观）
func ReverseInteger_COOLFUN(x int) int {
	if math.MaxInt32 < x ||  math.MinInt32 > x{
		return 0
	}

	var smallerThanZero = false
	var newX = x
	if x <  0 {
		smallerThanZero  = true
		newX =  newX * -1
	}

	newInt := 0
	for newX > 0 {
		remain := newX  % 10
		newInt  *= 10
		newInt += remain
		newX /= 10
	}

	if math.MaxInt32 < newInt ||  math.MinInt32 > newInt{
		return 0
	}
	if smallerThanZero {
		return newInt * -1
	}
	return newInt
}

//ReverseInteger 超过范围返回0
//若是负数，保存一个布尔变量，返回结果前作转换
//将数字转为字符串后，遍历字符组成新的字符串后转为int
func ReverseInteger(x int) int {

	if math.MaxInt32 < x ||  math.MinInt32 > x{
		return 0
	}

	var smallerThanZero = false
	var newInt = x
	if x <  0 {
		smallerThanZero  = true
		newInt =  newInt * -1
	}

	xStr := strconv.Itoa(newInt)
	nStr := ""
	for  i:= len(xStr); i >0; i--{
		nStr += string(xStr[i -1])
	}
	result,err :=   strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error conventing string to int")
	}
	if math.MaxInt32 < result ||  math.MinInt32 > result{
		return 0
	}
	if smallerThanZero {
		return result * -1
	}
	return result
}
