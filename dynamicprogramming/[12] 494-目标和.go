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

// todo
// 好的，变成背包问题的标准形式：
//
//有一个背包，容量为 sum，现在给你 N 个物品，第 i 个物品的重量为 nums[i - 1]（注意 1 <= i <= N），每个物品只有一个，请问你有几种不同的方法能够恰好装满这个背包？
//
//现在，这就是一个正宗的动态规划问题了，下面按照我们一直强调的动态规划套路走流程：
//
//第一步要明确两点，「状态」和「选择」。
//
//对于背包问题，这个都是一样的，状态就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」。
//
//第二步要明确 dp 数组的定义。
//
//按照背包问题的套路，可以给出如下定义：
//
//dp[i][j] = x 表示，若只在前 i 个物品中选择，若当前背包的容量为 j，则最多有 x 种方法可以恰好装满背包。
//
//翻译成我们探讨的子集问题就是，若只在 nums 的前 i 个元素中选择，若目标和为 j，则最多有 x 种方法划分子集。
//
//根据这个定义，显然 dp[0][..] = 0，因为没有物品的话，根本没办法装背包；dp[..][0] = 1，因为如果背包的最大载重为 0，「什么都不装」就是唯一的一种装法。
//
//我们所求的答案就是 dp[N][sum]，即使用所有 N 个物品，有几种方法可以装满容量为 sum 的背包。
//
//第三步，根据「选择」，思考状态转移的逻辑。
//
//回想刚才的 dp 数组含义，可以根据「选择」对 dp[i][j] 得到以下状态转移：
//
//如果不把 nums[i] 算入子集，或者说你不把这第 i 个物品装入背包，那么恰好装满背包的方法数就取决于上一个状态 dp[i-1][j]，继承之前的结果。
//
//如果把 nums[i] 算入子集，或者说你把这第 i 个物品装入了背包，那么只要看前 i - 1 个物品有几种方法可以装满 j - nums[i-1] 的重量就行了，所以取决于状态 dp[i-1][j-nums[i-1]]。
//
//PS：注意我们说的 i 是从 1 开始算的，而数组 nums 的索引时从 0 开始算的，所以 nums[i-1] 代表的是第 i 个物品的重量，j - nums[i-1] 就是背包装入物品 i 之后还剩下的容量。
//
//由于 dp[i][j] 为装满背包的总方法数，所以应该以上两种选择的结果求和，得到状态转移方程：
//
//代码语言：javascript代码运行次数：0
//运行
//AI代码解释
//dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]];
//然后，根据状态转移方程写出动态规划算法：
//
//代码语言：javascript代码运行次数：0
//运行
//AI代码解释
///* 计算 nums 中有几个子集的和为 sum */
//int subsets(int[] nums, int sum) {
//    int n = nums.length;
//    int[][] dp = new int[n + 1][sum + 1];
//    // base case
//    for (int i = 0; i <= n; i++) {
//        dp[i][0] = 1;
//    }
//
//    for (int i = 1; i <= n; i++) {
//        for (int j = 0; j <= sum; j++) {
//            if (j >= nums[i-1]) {
//                // 两种选择的结果之和
//                dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]];
//            } else {
//                // 背包的空间不足，只能选择不装物品 i
//                dp[i][j] = dp[i-1][j];
//            }
//        }
//    }
//    return dp[n][sum];
//}
//然后，发现这个 dp[i][j] 只和前一行 dp[i-1][..] 有关，那么肯定可以优化成一维 dp：
//
//代码语言：javascript代码运行次数：0
//运行
//AI代码解释
///* 计算 nums 中有几个子集的和为 sum */
//int subsets(int[] nums, int sum) {
//    int n = nums.length;
//    int[] dp = new int[sum + 1];
//    // base case
//    dp[0] = 1;
//
//    for (int i = 1; i <= n; i++) {
//        // j 要从后往前遍历
//        for (int j = sum; j >= 0; j--) {
//            // 状态转移方程
//            if (j >= nums[i-1]) {
//                dp[j] = dp[j] + dp[j-nums[i-1]];
//            } else {
//                dp[j] = dp[j];
//            }
//        }
//    }
//    return dp[sum];
//}
//对照二维 dp，只要把 dp 数组的第一个维度全都去掉就行了，唯一的区别就是这里的 j 要从后往前遍历，原因如下：
//
//因为二维压缩到一维的根本原理是，dp[j] 和 dp[j-nums[i-1]] 还没被新结果覆盖的时候，相当于二维 dp 中的 dp[i-1][j] 和 dp[i-1][j-nums[i-1]]。
//
//那么，我们就要做到：在计算新的 dp[j] 的时候，dp[j] 和 dp[j-nums[i-1]] 还是上一轮外层 for 循环的结果。
//
//如果你从前往后遍历一维 dp 数组，dp[j] 显然是没问题的，但是 dp[j-nums[i-1]] 已经不是上一轮外层 for 循环的结果了，这里就会使用错误的状态，当然得不到正确的答案。
//
//现在，这道题算是彻底解决了。
//
//总结一下，回溯算法虽好，但是复杂度高，即便消除一些冗余计算，也只是「剪枝」，没有本质的改进。而动态规划就比较玄学了，经过各种改造，从一个加减法问题变成子集问题，又变成背包问题，经过各种套路写出解法，又搞出状态压缩，还得反向遍历。
//
//现在搞得我都忘了自己是来干嘛的了。嗯，这也许就是动态规划的魅力吧。
