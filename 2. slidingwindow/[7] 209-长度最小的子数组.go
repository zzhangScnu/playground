package slidingwindow

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其
// 长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组[4,3]是该条件下的长度最小的子数组。
//
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
//
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 提示：
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
//
// 进阶：
//
// 如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
func minSubArrayLen(target int, nums []int) int {
	var i, val int
	res := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		val += nums[j]
		for val >= target {
			if j-i+1 < res {
				res = j - i + 1
			}
			val -= nums[i]
			i++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

/**
思路：
维护一个滑动窗口，i&j分别为其左边界&右边界。
- 不断扩展右边界，并累加子数组之和；
   1. 将右侧元素纳入滑动窗口：val += nums[j]；
   2. 扩展右边界：for j := 0; j < len(nums); j++
- 当子数组之和res满足约束时，即res >= target时，试图在该条件下尽量使得子数组长度最小：
   1. 即不断更新结果：res = min(res, j-i+1)；
   2. 将左侧元素排除出滑动窗口：val -= nums[i]；
   3. 缩小左边界：i++
可以看出，扩展&缩小是一对对称操作。

一些初始值：
慢指针&快指针：0
结果：数组长度+1，用来判断是否有被更新过、有无满足约束条件的场景
窗口区间：更新结果时，快慢指针都还未自增，所以是左闭右闭区间。
*/
