package sort

func quickSort(nums []int) []int {
	doQuickSort(nums, 0, len(nums)-1)
	return nums
}

func doQuickSort(nums []int, beginIdx, endIdx int) {
	if beginIdx < 0 || endIdx > len(nums)-1 {
		return
	}
	if beginIdx >= endIdx {
		return
	}
	pivot := partition(nums, beginIdx, endIdx)
	doQuickSort(nums, beginIdx, pivot-1)
	doQuickSort(nums, pivot+1, endIdx)
}

func partition(nums []int, beginIdx, endIdx int) int {
	pivot := nums[endIdx]
	slow, fast := beginIdx, beginIdx
	for ; fast < endIdx; fast++ {
		if nums[fast] < pivot {
			swap(nums, slow, fast)
			slow++
		}
	}
	swap(nums, slow, endIdx)
	return slow
}

/**
原地排序算法，无需额外空间复杂度。
平均时间复杂度为O(nlogn)。

合并排序算法：从下往上，先分解，再合并
快速排序算法：从上往下，先分区，再分解
*/
