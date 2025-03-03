package dynamicprogramming

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
// 被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
// 示例 1：
//
// 输入：[1,2,3,1]
// 输出：4
// 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//
//	偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
// 输入：[2,7,9,3,1]
// 输出：12
// 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//
//	偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
// 提示：
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

/**
DP数组及下标含义：
- i：偷/不偷第i个房屋；
- dp[i]：偷/不偷第i个房屋时，能偷窃的最大金额。

递推公式：
dp[i] = max(dp[i-1], dp[i-2] + nums[i])
dp[i-2] + nums[i]：偷当前房屋，因为防盗系统限制不能偷连续房屋，所以此时只能考虑第i-2间房屋的偷盗金额；
dp[i-1]：不偷当前房屋，考虑第i-1间房屋的偷盗金额。

注意，由dp数组定义，dp[i-1]和dp[i-2]是可能偷，也可能不偷的结果。
特别是dp[i-1]，并不是说不偷i，就一定会偷i-1。因为这样不一定是最佳策略。

初始化：
dp[0] = nums[0], dp[1] = max(dp[0], nums[1])

推导方向：从左到右

要注意因为递推公式有i-2的操作，所以需要保证输入数组长度>=2，防止越界。
*/
