package backtracking

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
// 只使用数字1到9
// 每个数字 最多使用一次
//
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
// 示例 1:
//
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 解释:
// 1 + 2 + 4 = 7
// 没有其他符合的组合了。
//
// 示例 2:
//
// 输入: k = 3, n = 9
// 输出: [[1,2,6], [1,3,5], [2,3,4]]
// 解释:
// 1 + 2 + 6 = 9
// 1 + 3 + 5 = 9
// 2 + 3 + 4 = 9
// 没有其他符合的组合了。
//
// 示例 3:
//
// 输入: k = 4, n = 1
// 输出: []
// 解释: 不存在有效的组合。
// 在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
// 提示:
//
// 2 <= k <= 9
// 1 <= n <= 60
func combinationSum3(k int, n int) [][]int {
	combination = []int{}
	combinations = [][]int{}
	doCombinationSum3(1, k, n)
	return combinations
}

func doCombinationSum3(beginIdx int, k int, remainTarget int) {
	if remainTarget < 0 {
		return
	}
	if len(combination) == k && remainTarget == 0 {
		res := make([]int, len(combination))
		copy(res, combination)
		combinations = append(combinations, res)
		return
	}
	for i := beginIdx; i <= 9-(k-len(combination))+1; i++ {
		combination = append(combination, i)
		doCombinationSum3(i+1, k, remainTarget-i)
		combination = combination[:len(combination)-1]
	}
}

/**
整体思路+剪枝
*/
