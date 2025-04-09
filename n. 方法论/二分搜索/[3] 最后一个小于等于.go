package 二分搜索

/*
nums = [5, 5, 7, 8, 8, 10]，target = 6，res = 1

思路类似于"寻找最后一个位置"，
唯一不同的是，在nums[mid] < target && nums[mid] == target时，
均需要判断nums[mid]是否 <= target的最后一个元素，若是则直接返回：
  - 若mid已经遍历到数组末尾；
  - 若mid右侧的元素 > target；

若否，推进low，在[mid + 1, high]中继续寻找。
*/
func searchLastLE(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			if mid == n-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
