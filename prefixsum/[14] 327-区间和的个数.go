package prefixsum

// 给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和
// upper）之内的 区间和的个数 。
//
// 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
//
// 示例 1：
//
// 输入：nums = [-2,5,-1], lower = -2, upper = 2
// 输出：3
// 解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
//
// 示例 2：
//
// 输入：nums = [0], lower = 0, upper = 0
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// -10⁵ <= lower <= upper <= 10⁵
// 题目数据保证答案是一个 32 位 的整数

var rangeSumCount int

var tempInCountRangeSum []int

func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	rangeSumCount, tempInCountRangeSum = 0, make([]int, len(preSum))
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	sortInCountRangeSum(preSum, lower, upper, 0, len(preSum)-1)
	return rangeSumCount
}

func sortInCountRangeSum(preSum []int, lower, upper int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)>>1
	sortInCountRangeSum(preSum, lower, upper, low, mid)
	sortInCountRangeSum(preSum, lower, upper, mid+1, high)
	mergeInCountRangeSum(preSum, lower, upper, low, mid, high)
}

func mergeInCountRangeSum(preSum []int, lower, upper int, low, mid, high int) {
	left, right := mid+1, mid+1
	for cur := low; cur <= mid; cur++ {
		for left <= high && preSum[left]-preSum[cur] < lower {
			left++
		}
		for right <= high && preSum[right]-preSum[cur] <= upper {
			right++
		}
		rangeSumCount += right - left
	}
	for i := low; i <= high; i++ {
		tempInCountRangeSum[i] = preSum[i]
	}
	i, j := low, mid+1
	for cur := low; cur <= high; cur++ {
		if i == mid+1 {
			preSum[cur] = tempInCountRangeSum[j]
			j++
		} else if j == high+1 {
			preSum[cur] = tempInCountRangeSum[i]
			i++
		} else if tempInCountRangeSum[i] <= tempInCountRangeSum[j] {
			preSum[cur] = tempInCountRangeSum[i]
			i++
		} else {
			preSum[cur] = tempInCountRangeSum[j]
			j++
		}
	}
}
