package dynamicprogramming

// KnapsackV
// 有N种物品和一个容量为W的背包。
// 第i种物品最多有Ni件可用，每件耗费的空间是Wi，价值是Vi。
// 求解将哪些物品装入背包可使这些物品的耗费的空间总和不超过背包容量，且价值总和最大。
func KnapsackV(W int, nums []int, weights []int, values []int) int {
	dp := make([]int, W+1)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := W; j >= weights[i]; j-- {
			for k := 1; k <= nums[i] && j >= weights[i]*k; k++ {
				dp[j] = max(dp[j], dp[j-weights[i]*k]+values[i]*k)
			}
		}
	}
	return dp[W]
}

func KnapsackVI(W int, nums []int, weights []int, values []int) int {
	for i, count := range nums {
		for j := 1; j < count; j++ {
			weights = append(weights, weights[i])
			values = append(values, values[i])
		}
	}
	dp := make([]int, W+1)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := W; j >= weights[i]; j-- {
			if j >= weights[i] {
				dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
			}
		}
	}
	return dp[W]
}

/**
多重背包实际上跟0/1背包很像，只是每件物品的数量从1件变成了多件。
实际上可以将n件展开为n个1件，本质上还是对每件物品取/不取，且只能取一次的0/1背包问题。

方法1，KnapsackV：
在原有的物品-背包遍历基础上，增加对每种物品的每件的遍历。
如果对物品i取1件，则公式为dp[j] = max(dp[j], dp[j-weights[i]*1]+values[i]*1)，
即dp[j] = max(dp[j], dp[j-weights[i]]+values[i])，就是0/1背包的递推公式。
如果是[2, nums[i]]件，则需空出的背包容量和增加的物品价值都需要累加。

一开始写成了
for k := 1; k <= nums[i] && j >= weights[i]*k; k++ {
	if j >= weights[i]*k {
		dp[j] = max(dp[j], dp[j-weights[i]*k]+values[i]*k)
	}
}
实际上，如果背包已经不足以装下k件物品i时，也不需要继续向后尝试了，一定装不下。所以作为结束条件即可。
for k := 1; k <= nums[i] && j >= weights[i]*k; k++

方法2，KnapsackVI：
根据nums数组，将weights和values打平成每种物品1件的形式，那么递推公式就跟0/1背包完全一致了。
但是对数组的赋值、扩容等操作可能会较为耗时且占空间。
*/
