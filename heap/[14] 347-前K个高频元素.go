package heap

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
// 示例 1:
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
//
// 示例 2:
//
// 输入: nums = [1], k = 1
// 输出: [1]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}
	minHeap := NewMinHeap()
	for num, cnt := range cntMap {
		minHeap.Insert(num, cnt)
		if minHeap.Size() > k {
			minHeap.ExtractMin()
		}
	}
	var res []int
	for !minHeap.IsEmpty() {
		res = append(res, minHeap.ExtractMin().Num)
	}
	return res
}

/**
实现了一个存储结构体的最小堆。
求前K个高频元素，可以维护一个大小为K的最小堆，堆顶是第K个高频元素，孩子节点存储的都是比它更高频的元素。

for num, cnt := range cntMap {
	minHeap.Insert(num, cnt)
	if minHeap.Size() > k {
		minHeap.ExtractMin()
	}
}
实际上，堆的大小是K+1，当堆满时，其中的K+1个元素已按出现次数降序排好。这时候移除堆顶元素，即移除第K+1个最高频元素，
重新堆化后就剩下K个最高频元素。
也可以维护大小为K的堆，当堆满时，若有新元素要入堆，则与堆顶元素出现频率判断，如果更小，则跳过不处理。
for num, cnt := range cntMap {
	if minHeap.Size() < k {
		minHeap.Insert(num, cnt)
		continue
	}
	if cnt < minHeap.PeekMin().Cnt {
		continue
	}
	minHeap.ReplaceMin(num, cnt)
}
*/
