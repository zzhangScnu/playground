package array

import "slices"

// 给你一个整数数组 nums 和一个整数 k 。
//
// 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
//
// 返回你可以对数组执行的最大操作数。
//
// 示例 1：
//
// 输入：nums = [1,2,3,4], k = 5
// 输出：2
// 解释：开始时 nums = [1,2,3,4]：
// - 移出 1 和 4 ，之后 nums = [2,3]
// - 移出 2 和 3 ，之后 nums = []
// 不再有和为 5 的数对，因此最多执行 2 次操作。
//
// 示例 2：
//
// 输入：nums = [3,1,3,4,3], k = 6
// 输出：1
// 解释：开始时 nums = [3,1,3,4,3]：
// - 移出前两个 3 ，之后nums = [1,4,3]
// 不再有和为 6 的数对，因此最多执行 1 次操作。
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
// 1 <= k <= 10⁹
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	res, count := 0, 0
	for left < right {
		res = nums[left] + nums[right]
		if res < k {
			left++
		} else if res > k {
			right--
		} else {
			left++
			right--
			count++
		}
	}
	return count
}

/**
思路：
为了凑出某个数，本质上是取1小1大，
所以先进行排序，从两边向中间逼近，每当遇到和为 k 的情况，就累加计数值。
如果始终没遇到，那就是凑不出来。
*/
