package 二分搜索

/*
在基本框架的基础上，处理nums[mid] == target分支时不是直接返回索引，
而是继续往左去找。若左边还有元素等于target，则在[left, mid - 1]区间内，收缩right。
本质上跟逐个向左遍历的线性查找是一样的，只是利用数组有序的特性将搜索空间每次砍半，从而达到O(logN)的时间复杂度。

为什么用left来判断是否找到目标值？
nums[mid] >= target -> right = mid - 1 -> right不断向左推进，nums[right + 1 ... N-1] >= nums[right]
nums[mid] < target -> left = mid + 1 -> left不断向右推进，nums[0 ... left - 1] < nums[left]
循环结束时，left > right，即left == right + 1，[right, left]区间内没有元素。
此时left左边的元素(不包含left)均比target小，而right右边的元素(不包含right)均比target大或相等，故right + 1即left指向的元素即可能为第一个target出现的位置。
如[5, 7, 7, 8, 8, 10]，结束时right = 2，left = 3。

分情况讨论：
1. left == N：数组越界，没有找到目标值，此时应返回-1；
2. left < 0：由于left只会基于mid递增，不会出现小于0的情况；
3. left在[0, N]之间：若nums[left] != target，没有找到目标值，此时应返回-1。
*/
func searchLeftRange(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left
}

/*
这种写法逻辑更清晰，
在基本框架的基础上，处理nums[mid] == target分支时不是直接返回索引，
而是先判断：
- 若mid == 0：说明已找到等于target子序列的起始位置，直接返回；
- 若mid > 0：若mid左边的元素 != target，说明mid指向target子序列的起始位置，直接返回；
如果不是这两种情况，则说明没有找到target出现的第一个位置，需继续往mid左边[left, mid - 1]范围寻找，即收缩right。
*/
func searchLeftRangeII(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
