package sort

var temp []int

func mergeSortII(nums []int) []int {
	n := len(nums)
	temp = make([]int, n)
	sort(nums, 0, n-1)
	return nums
}

func sort(nums []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)>>1
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	mergeII(nums, lo, mid, hi)
}

func mergeII(nums []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}
	i, j := lo, mid+1
	for cur := lo; cur <= hi; cur++ {
		if i == mid+1 {
			nums[cur] = temp[j]
			j++
		} else if j == hi+1 {
			nums[cur] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			nums[cur] = temp[i]
			i++
		} else {
			nums[cur] = temp[j]
			j++
		}
	}
}

/**
归并排序思路2-原地排序：
不像思路1-分解为子数组，通过截断原数组、重组新数组来实现，需要维护额外空间，而是通过全局变量+原地排序实现。
- mergeII的目的：将nums中[lo, hi]索引区间中的元素原地排序。其中nums[lo, mid]，nums[mid + 1, hi]已通过子问题分解和处理达到有序状态。
- 步骤详解：
	- 将nums[lo, hi]复制到全局辅助数组temp，防止对nums赋值时污染原始排序数据；
	- 将i指向nums[lo, mid]的起点lo，j指向nums[mid + 1, hi]的起点mid + 1；
	- 对索引范围[lo, hi]进行遍历，依次对nums中该范围内的元素进行原地排序：
		- 若此时i == mid + 1，说明已经遍历完nums[lo, mid]，此时需要将nums[mid + 1, hi]中的元素赋予nums的当前位置：nums[cur] = temp[j]，j++；
		- 若此时j == hi + 1，说明已经遍历完nums[mid + 1, hi]，此时需要将nums[lo, mid]中的元素赋予nums的当前位置：nums[cur] = temp[i]，i++；
		- 否则说明左 / 右数组都仍有元素剩余，需要处理，此时比较两边元素的大小，将较小的放入nums的当前位置。
	- 循环此过程，当lo == 0，hi == len(nums) - 1时，意味着nums中的元素均自底向上排序完毕，此时nums有序。
*/
