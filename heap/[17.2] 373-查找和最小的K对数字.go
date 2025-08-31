package heap

import "container/heap"

package heap

import "container/heap"

// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
// 示例 1:
//
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//
//	[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
// 示例 2:
//
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//
//	[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
// 提示:
//
// 1 <= nums1.length, nums2.length <= 10⁵
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1 和 nums2 均为 升序排列
//
// 1 <= k <= 10⁴
// k <= nums1.length * nums2.length
func kSmallestPairsII(nums1 []int, nums2 []int, k int) [][]int {
	minHeap := &IndexPairHeap{}
	heap.Init(minHeap)
	for i := range nums1 {
		heap.Push(minHeap, []int{i, 0})
	}
	var res [][]int
	for minHeap.Len() > 0 && k > 0 {
		cur := heap.Pop(minHeap).([]int)
		res = append(res, cur)
		if cur[1]+1 < len(nums2) {
			heap.Push(minHeap, []int{cur[0], cur[1] + 1})
		}
	}
	return res
}

type IndexPairHeap struct {
	nums1, nums2 []int
	indexes      [][]int
}

func (s IndexPairHeap) Len() int {
	return len(s.indexes)
}

func (s IndexPairHeap) Less(i, j int) bool {
	return s.nums1[s.indexes[i][0]]+s.nums2[s.indexes[i][1]] < s.nums1[s.indexes[j][0]]+s.nums2[s.indexes[j][1]]
}

func (s IndexPairHeap) Swap(i, j int) {
	s.indexes[i], s.indexes[j] = s.indexes[j], s.indexes[i]
}

func (s *IndexPairHeap) Push(x any) {
	(*s).indexes = append((*s).indexes, x.([]int))
}

func (s *IndexPairHeap) Pop() any {
	res := (*s).indexes[s.Len()-1]
	(*s).indexes = (*s).indexes[:s.Len()-1]
	return res
}
