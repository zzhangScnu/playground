package ___滑动窗口

// 给你一个整数数组 nums 和一个整数 k ，找出 nums 中和至少为 k 的 最短非空子数组 ，并返回该子数组的长度。如果不存在这样的 子数组 ，返回
// -1 。
//
// 子数组 是数组中 连续 的一部分。
//
// 示例 1：
//
// 输入：nums = [1], k = 1
// 输出：1
//
// 示例 2：
//
// 输入：nums = [1,2], k = 4
// 输出：-1
//
// 示例 3：
//
// 输入：nums = [2,-1,2], k = 3
// 输出：3
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁵ <= nums[i] <= 10⁵
// 1 <= k <= 10⁹

var MaxLen = 100_001

type MonotonicDeque [][]int

func (m *MonotonicDeque) Len() int {
	return len(*m)
}

func (m *MonotonicDeque) Push(x []int) {
	for len(*m) > 0 && (*m)[len(*m)-1][0] > x[0] {
		*m = (*m)[:len(*m)-1]
	}
	*m = append(*m, x)
}

func (m *MonotonicDeque) Pop() []int {
	if len(*m) > 0 {
		val := (*m)[0]
		*m = (*m)[1:]
		return val
	}
	return nil
}

func (m *MonotonicDeque) Max() []int {
	if len(*m) == 0 {
		return nil
	}
	return (*m)[0]
}

func shortestSubarray(nums []int, k int) int {
	res := MaxLen
	deque := &MonotonicDeque{}
	curSum := 0
	for index, num := range nums {
		curSum += num
		if curSum >= k {
			res = min(res, index+1)
		}
		for deque.Len() > 0 && curSum-deque.Max()[0] >= k {
			res = min(res, index-deque.Pop()[1])
		}
		deque.Push([]int{curSum, index})
	}
	if res == MaxLen {
		return -1
	}
	return res
}
