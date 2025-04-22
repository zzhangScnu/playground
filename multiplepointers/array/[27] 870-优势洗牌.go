package array

import (
	"container/heap"
	"slices"
)

// 给定两个长度相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的
// 数目来描述。
//
// 返回 nums1 的 任意 排列，使其相对于 nums2 的优势最大化。
//
// 示例 1：
//
// 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
// 输出：[2,11,7,15]
//
// 示例 2：
//
// 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
// 输出：[24,32,8,12]
//
// 提示：
//
// 1 <= nums1.length <= 10⁵
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 10⁹

type MaxHeap [][]int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i][1] > m[j][1]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *MaxHeap) Pop() any {
	val := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return val
}

func advantageCount(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for index, num := range nums2 {
		heap.Push(maxHeap, []int{index, num})
	}
	res := make([]int, len(nums1))
	left, right := 0, len(nums1)-1
	for maxHeap.Len() > 0 {
		candidate := heap.Pop(maxHeap).([]int)
		if nums1[right] > candidate[1] {
			res[candidate[0]] = nums1[right]
			right--
		} else {
			res[candidate[0]] = nums1[left]
			left++
		}
	}
	return res
}

/**
思路：
本题求的是最大优势，即不强求nums1中每个元素都能大于nums2中对应位置的元素，只求战胜数量最多的排列即可。

参考田忌赛马，如果可战胜，则派出上等马，为自己赢得一局；如果不可战胜，则用派出下等马，消耗对方战力较强的一名选手。
所以需要动态维护对方选手，每次从中挑选战力最强的选手candidate，与己方进行比较；
而己方也需要前置升序排序。
先看己方队尾nums1[right]，即战力最强的选手，是否可战胜candidate。如果是，则派出一战，优势+1；

假设此时nums1[right] >> candidate，是否需要保留实力，选取可战胜candidate但战力排名较后的选手？
答：没必要，因为candidate此时已经是对方最强的选手，即使保留当前nums1[right]，后续也没有更强的对手需要面对。
所以这里无需使用贪心思想，直接出战即可。

如果无法战胜，则保留实力，派出己方队头nums1[left]，即战力最弱的选手送人头，消耗对方战力较强的一名选手。

反复执行此对战，直到对方选手全部出战完毕。

因为要保留nums2选手的出场顺序，即nums2中各元素的相对位置，故不能对nums2直接进行排序。
可借助最大堆来维护(nums2中索引，nums2中元素)。
其中nums2中元素用于比较；nums2中索引用于排布nums1中元素的最终出场顺序。
*/
