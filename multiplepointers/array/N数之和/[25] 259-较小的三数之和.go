package N数之和

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

/**
使用双向指针的前提是数组有序，故先排序。
依次固定三数中的第一个数num，基于此下探选取第二、第三个数，其和需 < target - num。

在第二、第三个数的选取中，从num的下一个位置start开始搜索。使用双向指针，通过指向的两数之和sum的与k的关系，来推动左 / 右指针的移动。
同时在满足sum < target - num约束时，收集结果。
注意，因为区间[low, high]递增，且nums[low] + nums[high] < target - num，故对于[low, high)中的所有元素x，均有x + nums[high] < target - num，
即区间[low, high]中，满足条件的二元组数量为high - low。

注意，第一个数的范围为[0, len(nums)-2)，因为要预留位置给第二、第三个数。

base case：
if len(nums) < 3 {
	return 0
}
*/
