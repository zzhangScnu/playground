package backtracking

import "slices"

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
// 示例 1：
//
// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
// 示例 2：
//
// 输入：nums = [0]
// 输出：[[],[0]]
//
// 提示：
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	var subset []int
	var subsets [][]int
	slices.Sort(nums)
	var doSubsetsWithDup func(beginIdx int)
	doSubsetsWithDup = func(beginIdx int) {
		subsets = append(subsets, append([]int{}, subset...))
		for i := beginIdx; i < len(nums); i++ {
			if i > beginIdx && nums[i-1] == nums[i] {
				continue
			}
			subset = append(subset, nums[i])
			doSubsetsWithDup(i + 1)
			subset = subset[:len(subset)-1]
		}
	}
	doSubsetsWithDup(0)
	return subsets
}

/**
去重：均需提前排序
方法1：i > beginIdx
方法2：used局部变量-某节点下的同一层
方法3：used全局变量-某节点下的子树
*/
