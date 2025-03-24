package array

import "slices"

// 给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
//
// 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
//
// 以数组形式返回答案。
//
// 示例 1：
//
// 输入：nums = [8,1,2,2,3]
// 输出：[4,0,1,1,3]
// 解释：
// 对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
// 对于 nums[1]=1 不存在比它小的数字。
// 对于 nums[2]=2 存在一个比它小的数字：（1）。
// 对于 nums[3]=2 存在一个比它小的数字：（1）。
// 对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
//
// 示例 2：
//
// 输入：nums = [6,5,4,8]
// 输出：[2,1,0,3]
//
// 示例 3：
//
// 输入：nums = [7,7,7,7]
// 输出：[0,0,0,0]
//
// 提示：
//
// 2 <= nums.length <= 500
// 0 <= nums[i] <= 100
func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	slices.Sort(sorted)
	numPosMap := make(map[int]int)
	for i := n - 1; i >= 0; i-- {
		numPosMap[sorted[i]] = i
	}
	res := make([]int, n)
	for i, num := range nums {
		res[i] = numPosMap[num]
	}
	return res
}

func smallerNumbersThanCurrentII(nums []int) []int {
	counter := make([]int, 101)
	for _, num := range nums {
		counter[num]++
	}
	prefix := make([]int, 101)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + counter[i-1]
	}
	for i, num := range nums {
		nums[i] = prefix[num]
	}
	return nums
}

/*
思路一：排序 + HashMap
1. 对原始数组进行升序排序；
2. 倒序遍历排序后的数组，用map记录每个数字的下标。目的：
  - 记录升序数组中的【元素：索引】映射，本质上是记录有多少元素小于当前元素；
  - 倒序：最后一次出现的数字，一定相对位置最靠左的，满足了【小于】条件，即更左的元素都更小，而不是【小于等于】。
注意对数组进行copy前，一定要记得初始化目标数组长度。
其实本质也是一种前缀和的做法。

思路二：前缀和数组
1. 由于题目给定数据范围，可以初始化一个数组作为哈希表；
2. 记录每个元素出现次数，由于前缀和数组下标有序，结果天然就是升序的。
*/
