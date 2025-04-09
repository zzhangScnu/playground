package 二分搜索

/*
nums = [5, 5, 7, 8, 8, 10]，target = 6，res = 2

思路类似于"寻找第一个位置"，
唯一不同的是，在nums[mid] > target && nums[mid] == target时，
均需要判断nums[mid]是否 >= target的第一个元素，若是则直接返回：
  - 若mid已经遍历到数组起始；
  - 若mid左侧的元素 < target；

若否，推进high，在[low, mid - 1]中继续寻找。
*/
func searchFirstGE(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

/*
这种写法体现了返回值i有不同的语义：
- 第一个大于等于target的位置为i；
- 有i个元素比target小；
- target插入nums中的位置为i。
*/
func searchFirstGEII(nums []int, num int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
