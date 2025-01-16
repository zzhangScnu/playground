package sort

import "math"

func mergeSort(nums []int) []int {
	return doMergeSort(nums, 0, len(nums)-1)
}

func doMergeSort(nums []int, beginIdx, endIdx int) []int {
	if beginIdx == endIdx {
		return []int{nums[beginIdx]}
	}
	mid := beginIdx + (endIdx-beginIdx)/2
	nums1 := doMergeSort(nums, beginIdx, mid)
	nums2 := doMergeSort(nums, mid+1, endIdx)
	return merge(nums1, nums2)
}

func merge(nums1, nums2 []int) []int {
	INF := math.MaxInt64
	nums1, nums2 = append(nums1, INF), append(nums2, INF)
	var i, j int
	var res []int
	for nums1[i] != INF || nums2[j] != INF {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	return res
}

/**
归并排序：
- 思路：分而治之
   将问题以同样形式分解为子问题，在到达base case(触底)时，开始回溯(反弹)，逐一解决拆分后的子问题。
- 实现：递归
	使用游标来控制分解&回溯的子问题边界。

哨兵解决边界问题：
普通实现方式，需要在较短数组处理完成后，将较长数组剩余元素逐一拷贝到结果集，需要处理数组越界情况；
但通过在原始数组末尾增加哨兵的形式，可以简化处理逻辑。

需要O(n)的额外空间复杂度，在每次合并过程中作为临时结果集。
平均时间复杂度为O(nlogn)。
*/
