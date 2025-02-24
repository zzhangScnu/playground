package array

// 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
// nums[i] 的元素的数量。
//
// 示例 1：
//
// 输入：nums = [5,2,6,1]
// 输出：[2,1,1,0]
// 解释：
// 5 的右侧有 2 个更小的元素 (2 和 1)
// 2 的右侧仅有 1 个更小的元素 (1)
// 6 的右侧有 1 个更小的元素 (1)
// 1 的右侧有 0 个更小的元素
//
// 示例 2：
//
// 输入：nums = [-1]
// 输出：[0]
//
// 示例 3：
//
// 输入：nums = [-1,-1]
// 输出：[0,0]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	var sorted []int
	for i := len(nums) - 1; i >= 0; i-- {
		index := findFarLeftIndex(sorted, nums[i])
		sorted = insert(sorted, index, nums[i])
		res[i] = index
	}
	return res
}

func findFarLeftIndex(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func insert(nums []int, index int, element int) []int {
	if len(nums) == 0 {
		return []int{element}
	}
	nums = append(nums, 0)
	copy(nums[index+1:], nums[index:])
	nums[index] = element
	return nums
}

