package backtracking

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
// 示例 1：
//
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：
// 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
// 7 也是一个候选， 7 = 7 。
// 仅有这两种组合。
//
// 示例 2：
//
// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3：
//
// 输入: candidates = [2], target = 1
// 输出: []
//
// 提示：
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// candidates 的所有元素 互不相同
// 1 <= target <= 40

func combinationSum(candidates []int, target int) [][]int {
	combination = []int{}
	combinations = [][]int{}
	doCombinationSum(0, candidates, target)
	return combinations
}

func doCombinationSum(beginIdx int, candidates []int, remainTarget int) {
	if remainTarget < 0 {
		return
	}
	if remainTarget == 0 {
		res := make([]int, len(combination))
		copy(res, combination)
		combinations = append(combinations, res)
		return
	}
	for i := beginIdx; i < len(candidates); i++ {
		combination = append(combination, candidates[i])
		doCombinationSum(i, candidates, remainTarget-candidates[i])
		combination = combination[:len(combination)-1]
	}
}

/**
无重复元素，候选元素可重复使用，结果集无重复。
结果集没有长度限制。

- 候选元素可重复使用：用doCombinationSum(i, ...) 来控制下一个元素的选取，仍然可以选到本元素；
  这里好容易写成beginIdx，需要特别注意不要手滑；
- 剪枝：可以先对候选元素进行整体排序，如果【单条路径当前的和 + 下一个候选元素 > 结果要求】，则无需考虑了。
*/
