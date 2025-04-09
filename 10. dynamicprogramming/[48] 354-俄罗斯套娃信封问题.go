package dynamicprogramming

import (
	"sort"
)

// 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
// 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出：3
// 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
// 示例 2：
//
// 输入：envelopes = [[1,1],[1,1],[1,1]]
// 输出：1
//
// 提示：
//
// 1 <= envelopes.length <= 10⁵
// envelopes[i].length == 2
// 1 <= wi, hi <= 10⁵
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	n := len(envelopes)
	var maxLen int
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}

/**
思路：
本质上是最长递增子序列，只不过是二维的。
将宽度升序排序，达到降维效果，只需对高度进行最长递增子序列的查找即可。
其长度即为最多可套信封的数量。

注意，更新maxLen的地方在i循环中，即处理完dp[i]后再更新。

为什么宽度相同的再按高度降序排序？
为了避免宽度相同的信封被重复选取。
例如[3, 4], [1, 2], [3, 5]，排序后为[1, 2], [3, 5], [3, 4]，
那么对高度序列查找最长递增子序列时，2 -> 5 -> 4x，4不会被选取，故规避了[3, 4]和[3, 5]同时被选取的情况。
*/
