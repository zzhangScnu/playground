package 方法论

/*
*
当目标值 < 区间[low, high]中间元素nums[mid]值时

	-> 目标值在mid左侧
	-> 缩小区间右边界
	-> high = mid - 1

当目标值 > 区间[low, high]中间元素nums[mid]值时

	-> 目标值在mid右侧
	-> 扩大区间左边界
	-> low = mid + 1

注意nums中的元素需保证无重复，在nums[mid] == target时，返回mid。

关于mid的计算：
mid := (low + high) / 2
如果在low和high的值都较大的情况下，计算容易溢出 -> 转换为mid := low + (high - low) / 2
将 / 2转换为位运算，加快计算速度 -> mid := low + (high - low) >> 1
*/
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (low+high)>>1
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
