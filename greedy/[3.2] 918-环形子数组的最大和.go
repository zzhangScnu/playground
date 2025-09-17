package greedy

// 给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
//
// 环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个
// 元素是 nums[(i - 1 + n) % n] 。
//
// 子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不
// 存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
//
// 示例 1：
//
// 输入：nums = [1,-2,3,-2]
// 输出：3
// 解释：从子数组 [3] 得到最大和 3
//
// 示例 2：
//
// 输入：nums = [5,-3,5]
// 输出：10
// 解释：从子数组 [5,5] 得到最大和 5 + 5 = 10
//
// 示例 3：
//
// 输入：nums = [3,-2,2,-3]
// 输出：3
// 解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3
//
// 提示：
//
// n == nums.length
// 1 <= n <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
func maxSubarraySumCircular(nums []int) int {
	total, curMin, globalMin, curMax, globalMax := 0, 0, nums[0], 0, nums[0]
	for _, num := range nums {
		total += num
		curMin = min(curMin+num, num)
		globalMin = min(globalMin, curMin)
		curMax = max(curMax+num, num)
		globalMax = max(globalMax, curMax)
	}
	if globalMax < 0 {
		return globalMax
	}
	return max(globalMax, total-globalMin)
}

/**
nums = [a b c d e]


环形子数组最大和 =
	max(
		中间部分子数组最大和，如 [b, c, d]
		头尾相连部分子数组最大和，如[a, e] -> 数组总和 - 中间部分子数组最小和
	)


求当前子数组的最小和：
curMin = min(curMin+num, num)
- curMin + num：将当前元素合并进当前子数组中
- num：从当前元素开始自成一派，重新开始建立子数组


初始化：
total, curMin, globalMin, curMax, globalMax := 0, 0, nums[0], 0, nums[0]
令 globalMin = globalMax = nums[0]，否则如果是全负数的场景，globalMin和globalMax都无法更新。


特殊处理：
若数组中均为负数，则globalMax最终为负数，此时最大和子数组即为【最大的单个元素】，此时存储在globalMax中。
if globalMax < 0 {
	return globalMax
}
*/
