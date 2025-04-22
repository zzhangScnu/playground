package array

// 给你一个整数数组 nums ，请计算数组的 中心下标 。
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
//
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
//
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
//
// 示例 1：
//
// 输入：nums = [1, 7, 3, 6, 5, 6]
// 输出：3
// 解释：
// 中心下标是 3 。
// 左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
// 右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
//
// 示例 2：
//
// 输入：nums = [1, 2, 3]
// 输出：-1
// 解释：
// 数组中不存在满足此条件的中心下标。
//
// 示例 3：
//
// 输入：nums = [2, 1, -1]
// 输出：0
// 解释：
// 中心下标是 0 。
// 左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
// 右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
//
// 提示：
//
// 1 <= nums.length <= 10⁴
// -1000 <= nums[i] <= 1000
func pivotIndex(nums []int) int {
	n := len(nums) + 1
	sum := make([]int, n)
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 0; i < n-1; i++ {
		left, right := sum[i], sum[n-1]-sum[i+1]
		if left == right {
			return i
		}
	}
	return -1
}

func pivotIndexII(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	var left int
	for i, num := range nums {
		if left == sum-left-num {
			return i
		}
		left += num
	}
	return -1
}

/**
思路：
前缀和数组。
先用O(n)统计前缀和，在基于每一个元素判断其左侧和右侧的和，看是否相等。
由于[2, 1, -1]这种case的存在，前缀和数组需要整体向右偏移一位。sum长度初始化为n+1，sum[0] = 0，表示下标0左侧的元素之和为0。

II是I的精简版，在I的基础上，将前缀和数组简化为滚动数组。
left = 遍历的每一个元素
right = sum - left - num
*/
