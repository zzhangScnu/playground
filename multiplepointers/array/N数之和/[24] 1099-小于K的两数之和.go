package N数之和

import (
	"slices"
)

/**
给你一个整数数组nums和整数k，返回最大和sum，
满足存在i<j使得nums[i]+nums[j] = sum且sum<k。
如果没有满足此等式的i，j存在，则返回-1。

示例1：
输入：nums=[34，23，1，24，75，33，54，8]， k = 60
输出：58
解释：
34和24相加得到58，58小于60，满足题意。

示例2：
输入：nums=[10，20，30]， k = 15
输出：-1
解释：
我们无法找到和小于15的两个元素。
*/

func twoSumLessThanK(nums []int, k int) int {
	slices.Sort(nums)
	res := -1
	sum, low, high := 0, 0, len(nums)-1
	for low < high {
		sum = nums[low] + nums[high]
		if sum < k {
			res = max(res, sum)
			low++
		} else {
			high--
		}
	}
	return res
}

/**
先排序，再使用双向指针，通过指向的两数之和的与k的关系，来推动左 / 右指针的移动。
同时在满足sum < k约束时，收集结果。
*/
