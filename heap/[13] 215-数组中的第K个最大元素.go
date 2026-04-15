package heap

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
	maxHeap := NewMaxHeap()
	// maxHeap.HeapifyByShiftUp(nums)
	maxHeap.HeapifyByShiftDown(nums)
	var res int
	for i := 0; i < k; i++ {
		res = maxHeap.ExtractMax()
	}
	return res
}

/**
思路：
1. 最小堆-大小为K：
较小的元素上浮到堆口，当元素数量超过容量K时，就将堆口元素弹出。（因为这些元素较小，而求解的是前K个最大元素）
当所有元素都处理完一遍之后，堆中留下的即为最大的K个元素，而堆顶元素是其中最小的元素，
即最终堆口的元素即为第K个最大元素。
2. 最大堆-大小为数组长度：
将数组中所有元素都入堆进行堆化。
再结合计数器，不断从堆口弹出元素，获取当前最大值，直至拿到第K个最大元素。
*/
