package easy

/*
1. Two Sum
Easy

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 */

//TwoSum_faster -- 遍历nums，假设索引为i的数字为v，target - v得到所需数字为v1，
//若能从numsMap从取得v1在nums中对应的索引INDEX，则返回[i,index]，
//否则将v和i存入numsMap中v为key,i为值
//只需要遍历一次
func TwoSum_faster(nums []int, target int) []int {
	numsMap := make(map[int]int)
	result := make([]int,0)
	for i,v := range nums{
		v1 := target - v
		index, ok := numsMap[v1]
		if ok == true {
			return append(result,i,index)
		}else {
			numsMap[v] = i
		}
	}
	return result
}

//TwoSum -- 遍历nums，假设索引为i的数字为v，target - v得到所需数字为v1，
//从i+1开始查找，找到值为v1的数字(假设其索引为a)则返回[i,a]
//可能需要多次遍历
func TwoSum(nums []int, target int) []int {
	for i,v := range nums {
		v1 := target - v
		for a:= i + 1; a < len(nums) ; a ++ {
			if v1 == nums[a] {
				return []int{i,a}
			}
		}
	}
	return []int{}
}