package heap

import "container/heap"

// KthLargest 设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
// 请实现 KthLargest 类：
//
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
//
// 示例 1：
//
// 输入： ["KthLargest", "add", "add", "add", "add", "add"] [[3, [4, 5, 8, 2]], [3]
// , [5], [10], [9], [4]]
//
// 输出：[null, 4, 5, 5, 8, 8]
//
// 解释：
//
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]); kthLargest.add(3); /
// / 返回 4 kthLargest.add(5); // 返回 5 kthLargest.add(10); // 返回 5 kthLargest.add(9);
// // 返回 8 kthLargest.add(4); // 返回 8
//
// 示例 2：
//
// 输入： ["KthLargest", "add", "add", "add", "add"] [[4, [7, 7, 7, 7, 8, 3]], [2],
// [10], [9], [9]]
//
// 输出：[null, 7, 7, 7, 8]
//
// 解释： KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
// kthLargest.add(2); // 返回 7
// kthLargest.add(10); // 返回 7
// kthLargest.add(9); // 返回 7
// kthLargest.add(9); // 返回 8
//
// 提示：
//
// 0 <= nums.length <= 10⁴
// 1 <= k <= nums.length + 1
// -10⁴ <= nums[i] <= 10⁴
// -10⁴ <= val <= 10⁴
// 最多调用 add 方法 10⁴ 次
type KthLargest struct {
	Capacity int
	*MinHeap
}

func KthLargestConstructor(k int, nums []int) KthLargest {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}
	return KthLargest{
		Capacity: k,
		MinHeap:  minHeap,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.MinHeap.Len() > this.Capacity {
		heap.Pop(this)
	}
	return (*this.MinHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
