package easy

/*
35. Search Insert Position
Easy

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

Example 1:
Input: nums = [1,3,5,6], target = 5
Output: 2

Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4

Example 4:
Input: nums = [1,3,5,6], target = 0
Output: 0

Example 5:
Input: nums = [1], target = 0
Output: 0

Constraints:
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
 */

func SearchInsert(nums []int, target int) int {

	n := len(nums)
	if n < 1 {
		return n
	}

	if n == 1 {
		if nums[0] < target {
			return 1
		}
		return 0
	}

	for i,v := range nums {
		if i < n-1 {
			vNext :=nums[i + 1]
			if v >= target  {
				return i
			}else if vNext >= target {
				return i+1
			}
		}
	}

	return n
}