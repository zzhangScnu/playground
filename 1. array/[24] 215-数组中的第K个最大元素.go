package array

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
// 提示：
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
func findKthLargest(nums []int, k int) int {
	beginIdx, endIdx := 0, len(nums)-1
	for {
		pivot := partition(nums, beginIdx, endIdx)
		if pivot+1 == k {
			return nums[pivot]
		}
		if pivot+1 > k {
			endIdx = pivot - 1
		} else {
			beginIdx = pivot + 1
		}
	}
}

func partition(nums []int, beginIdx, endIdx int) int {
	pivot := nums[endIdx]
	slow, fast := beginIdx, beginIdx
	for ; fast < endIdx; fast++ {
		if nums[fast] > pivot {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	nums[slow], nums[endIdx] = nums[endIdx], nums[slow]
	return slow
}

// 印象中以前的我貌似是用堆来做的额

/**
结合快排思想：本质又是抽象f(x)的二分查找
*/
