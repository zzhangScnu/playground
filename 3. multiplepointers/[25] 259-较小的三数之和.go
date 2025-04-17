package multiplepointers

import "slices"

/**
给定一个长度为n的整数数组和一个目标值target，寻找能够使条件nums[i] + nums[j] + nums [k]<target成立的三元组i，j，k个数(0<=i<j<k<n)。

示例1：
输入：nums
= [-2，0，1，3]， target = 2

输出：2

解释：因为一共有两个三元组满足累加和小于2：
[-2，0，1]
[-2，0，3]
*/

func threeSumSmaller(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	slices.Sort(nums)
	var res int
	for i := 0; i < len(nums)-2; i++ {
		res += twoSumSmaller(nums, i+1, target-nums[i])
	}
	return res
}

func twoSumSmaller(nums []int, start int, target int) int {
	var res int
	low, high := start, len(nums)-1
	for low < high {
		if nums[low]+nums[high] < target {
			res += high - low
			low++
		} else {
			high--
		}
	}
	return res
}
