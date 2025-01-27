package backtracking

import "slices"

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。
//
// 注意：解集不能包含重复的组合。
//
// 示例 1:
//
// 输入: candidates =[10,1,2,7,6,1,5], target =8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
//
// 示例 2:
//
// 输入: candidates =[2,5,2,1,2], target =5,
// 输出:
// [
// [1,2,2],
// [5]
// ]
//
// 提示:
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
func combinationSum2(candidates []int, target int) [][]int {
	combination = []int{}
	combinations = [][]int{}
	slices.Sort(candidates)
	used := make([]bool, len(candidates))
	doCombinationSum2(0, candidates, target, used)
	return combinations
}

func doCombinationSum2(beginIdx int, candidates []int, remainTarget int, used []bool) {
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
		if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
			continue
		}
		combination = append(combination, candidates[i])
		used[i] = true
		doCombinationSum2(i+1, candidates, remainTarget-candidates[i], used)
		used[i] = false
		combination = combination[:len(combination)-1]
	}
}

/**
有重复元素，候选元素不可重复使用，结果集无重复。

需要做去重。
思路：在树层上做去重。
当候选集为(1, 1`, 2)时，第一个分支选取1，该分支向下扩展时，可选取1`，这是相同值的不同元素的选取，是合法的；
但当从第一个分支回溯，弹出1后，在第二个分支选1`时，该分支向下扩展，选取2，此时一定会和第一个分支的某棵子树重复。
即(1, 1`)的情况下，第一个分支生成的子树一定包含第二个分支生成的子树，甚至将1`本身也包括在内。所以第二个分支实际是可以做剪枝的。

1. 对候选元素进行整体排序；
2. 引入used标记数组，记录候选集某下标对应的元素当前是否已被选取；
3. 判断条件：
	if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
		continue
	}
  【i > 0】：避免i-1数组越界；
  【candidates[i-1] == candidates[i] 】：数组非递减有序，相等表示当前元素跟之前处理的元素重复；
  【 !used[i-1]】：在树层中，前一个元素使用过。
                           意味着，是通过使用过前一个元素(used[i-1] = true)后，回溯(used[i-1] = false)来到当前节点，准备扩展新的一条分支的。
                           结合最开始思路，这种场景需要剪枝，continue即可。
  【used[i-1]】：在树枝中，前一个元素使用过。表示是树枝上的合法选择，使用1后再到下一层中使用1`的场景。
4. 这种判断条件实际也可以：
	if i > beginIdx && candidates[i-1] == candidates[i] {
		continue
	}
*/
