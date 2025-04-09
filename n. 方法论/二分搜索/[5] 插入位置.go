package 二分搜索

/*
在target元素可能不存在的nums中的场景下，
寻找target的插入位置可以转换为：寻找第一个大于等于target的元素位置。
*/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}
