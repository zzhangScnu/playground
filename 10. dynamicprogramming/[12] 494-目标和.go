package dynamicprogramming

import "math"

// 给你一个非负整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
// 示例 1：
//
// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
//
// 示例 2：
//
// 输入：nums = [1], target = 1
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum < int(math.Abs(float64(target))) {
		return 0
	}
	target = (sum + target) / 2
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1
	if target >= nums[0] {
		dp[0][nums[0]] = 1
	}
	var zeroCount float64
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zeroCount++
		}
		dp[i][0] = int(math.Pow(2, zeroCount))
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}

/**
二维数组解法：
将元素分为正负两个集合，可得：
positive + negative = sum
positive - negative = target
=>
positive = (sum + target)/2
即为找有多少种方法使得正数集合的和 == (sum + target)/2
每个元素值既为重量又为价值

DP数组及下标含义
- i：第0到第i个元素取/不取，[0, n-1]
- j：背包容量为j，[0, (sum + target)/2]
- dp[i][j]：第0到第i个元素依次取/不取，且背包容量为j时，一共有几种方式能够凑满j容量的包，即使得最终目标和为j

递推公式
不取当前元素i有几种方式：dp[i-1][j]
取当前元素i有几种方式：dp[i][j-nums[i]]
dp[i][j] = dp[i-1][j] + dp[i][j-nums[i]]

终止条件
if (sum+target)%2 == 1 || sum < int(math.Abs(float64(target))) {
	return 0
}
(sum+target)%2 == 1：无法整除，代入几个数字可验证无解；
sum < int(math.Abs(float64(target)))：非负候选集的总和都凑不到target的绝对值，表示不可能有结果。

遍历方向
从左往右、由上至下。

初始化
dp[0][0, (sum + target)/2]：只放物品0， 把容量为j的背包填满有几种方法。
只有背包容量为物品0的容量的时候，方法为1，正好装满。其他情况下，装不满或装不下。

dp[0][0]：将物品0装进容量为0的背包，默认只有1种方法，即为不装。如果物品0的重量为0，则为以下处理逻辑：
dp[0, n-1][0]：
- 如果物品0的容量 > 0，由于背包容量为0，使用默认零值即可；
- 否则要计算0可以组成0有几种方式。对每个0来说，有取/不取两种选择，则如果有n个0，就是n的2次方种方式。
*/

func findTargetSumWaysII(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum < int(math.Abs(float64(target))) {
		return 0
	}
	target = (sum + target) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}

/**
一维数组解法
消去i，直接在原一维数组上做计算，则递推公式变为：
dp[j] = dp[j] + dp[j-nums[i]]
换种方式理解，假设每次固定一个nums[i]：
已经有一个1(nums[i])的话，有dp[4]种方法凑成容量为5的背包。
已经有一个2(nums[i])的话，有dp[3]种方法凑成容量为5的背包。
已经有一个3(nums[i])的话，有dp[2]种方法凑成容量为5的背包
已经有一个4(nums[i])的话，有dp[1]种方法凑成容量为5的背包
已经有一个5(nums[i])的话，有dp[0]种方法凑成容量为5的背包
那么凑成dp[5]有多少种方法 == 把所有的dp[j - nums[i]]累加起来。

DP数组及下标含义
- j：容量为j的背包，[0, (sum + target)/2]]
- dp[j]：对于容量为j的背包，一共有dp[j]种方式装满

遍历方向：从左往右

初始化
dp[0]：1
*/
