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
候选集无重复，候选元素不可重复使用，结果集无重复。

- 不重复选取候选元素：用beginIdx+1作为树枝深度生长的入参，从而在下一层中作为for循环选取元素的初始位置。即不包括本层处理的元素；
- 全局变量：由于收集单条路径和收集所有结果集的变量都是全局变量，在每次主方法调用时都应清空；
- 回溯操作：
		combination = append(combination, i) // 考虑i，入i
		doCombinationSum3(i+1, k, remainTarget-i) // remainTarget-i是隐形回溯，如果单独用变量承接，也需要先减再加
		combination = combination[:len(combination)-1] // 不考虑i，弹出i，转而考虑i+1
- 剪枝操作：
  题目要求k个元素组合，所以当遍历到某些分支，单条路径的大小超过k，那就无需考虑了。
  所以可以通过for循环结束选取元素的位置来控制，当单条路径大小为m时，最多还可以在长度为n的候选集中选取多少个元素，即候选集的扫描区间范围，
  是[beginIdx, n-(k-m)+1]
*/
