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
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	minHeap := &SumIndexPairHeap{}
	heap.Init(minHeap)
	minHeap.Push([]int{nums1[0] + nums2[0], 0, 0})
	visited := make(map[[2]int]interface{})
	visited[[2]int{0, 0}] = struct{}{}
	var res [][]int
	for minHeap.Len() > 0 && k > 0 {
		cur := minHeap.Pop().([]int)
		i, j := cur[1], cur[2]
		res = append(res, []int{i, j})
		k--
		if _, ok := visited[[2]int{i + 1, j}]; !ok && i+1 < n {
			visited[[2]int{i + 1, j}] = struct{}{}
			minHeap.Push([]int{nums1[i+1] + nums2[j], i + 1, j})
		}
		if _, ok := visited[[2]int{i, j + 1}]; !ok && j+1 < m {
			visited[[2]int{i, j + 1}] = struct{}{}
			minHeap.Push([]int{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return res
}

type SumIndexPairHeap [][]int // sum, i from nums1, j from nums2

func (s SumIndexPairHeap) Len() int {
	return len(s)
}

func (s SumIndexPairHeap) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s SumIndexPairHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *SumIndexPairHeap) Push(x any) {
	*s = append(*s, x.([]int))
}

func (s *SumIndexPairHeap) Pop() any {
	res := (*s)[s.Len()-1]
	*s = (*s)[:s.Len()-1]
	return res
}
