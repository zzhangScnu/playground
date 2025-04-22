package slidingwindow

import (
	"container/heap"
)

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

type MinHeap [][]int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *MinHeap) Pop() any {
	val := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return val
}

func shortestSubarray(nums []int, k int) int {
	res, curSum := MaxLen, 0
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for index, num := range nums {
		curSum += num
		if curSum >= k {
			res = min(res, index+1)
		}
		for minHeap.Len() > 0 && curSum-(*minHeap)[0][0] >= k {
			res = min(res, index-heap.Pop(minHeap).([]int)[1])
		}
		heap.Push(minHeap, []int{curSum, index})
	}
	if res == MaxLen {
		return -1
	}
	return res
}

/**
思路：
暴力解法：
因为子数组必须连续，所以可以将假设当前索引指向位置为end，则可以将原数组nums分为几部分，
即[0, begin, end, len(nums)-1]。其中，
[0, begin)：遍历过，但不考虑纳入目标子数组的元素；
[begin, end]：目标子数组；
(end, len(nums)-1]：尚未遍历的元素。

故对于每一个元素nums[end]，可以遍历其之前的所有元素，找到不同的nums[begin]，使得[begin, end]满足和sum >= target且长度最小。
但这样的时间复杂度为O(n^2)。

前缀和 + 最小堆优化：
可以将时间复杂度从O(n^2)降为O(n)，但相应地空间复杂度从O(1)升为O(n)。
具体地，可以定义一个前缀和数组sums，其中sums[i]表示从nums[0]到nums[i]的元素和。
对于每一个元素nums[end]，可以不断地找到最小的sums[begin]，使得sums[end] - sums[begin] >= target。
为什么是最小的sums[begin]？
- 若sums[begin]为正数：最小值可以使得相减结果尽量大；
- 若sums[begin]为负数：减去负数实际上是加上其绝对值，故最小值可以使得相减结果尽量大。
两者都从尽量满足 >= target出发，体现了贪心思想。
此时可以更新最小子数组长度res = end - begin。

这个“找到最小的sums[begin]”，可以使用最小堆来实现。堆顶元素为最小的sums[i]。
同时，无需前置将前缀和一次性计算并维护在数组中，可以仅用一个变量curSum滚动计算并更新。

对于每一个元素nums[end]，可以将其加入当前位置前缀和curSum，与堆顶元素sums[begin]相减后再与target进行比较，如果满足条件，则将堆顶元素弹出，并尝试更新res。
否则跳过该元素、结束本次处理，推进end，将该前缀和留给下一个nums[end]使用。

可以看出curSum只是不断与nums[end]累加，并不真正减去前缀和。如果堆顶元素符合要求，则尝试下一个堆顶元素，即不断尝试前缀和更大的元素，试图找到更大的begin，即更小的子数组长度。

最后，将当前位置的前缀和压入最小堆中。注意这步需要放在最后，在当前位置计算完成且推进到下一位置后，当前的前缀和才能纳入考虑范围。
*/
