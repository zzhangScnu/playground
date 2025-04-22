package backtracking

// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
//
// 示例 1：
//
// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
//
// 示例 2：
//
// 输入：nums = [4,4,3,2,1]
// 输出：[[4,4]]
//
// 提示：
//
// 1 <= nums.length <= 15
// -100 <= nums[i] <= 100
func findSubsequences(nums []int) [][]int {
	var subsequence []int
	var subsequences [][]int
	var doFindSubsequences func(beginIdx int)
	doFindSubsequences = func(beginIdx int) {
		if len(subsequence) > 1 {
			subsequences = append(subsequences, append([]int{}, subsequence...))
		}
		used := make(map[int]interface{})
		for i := beginIdx; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			if len(subsequence) > 0 && subsequence[len(subsequence)-1] > nums[i] {
				continue
			}
			used[nums[i]] = true
			subsequence = append(subsequence, nums[i])
			doFindSubsequences(i + 1)
			subsequence = subsequence[:len(subsequence)-1]
		}
	}
	doFindSubsequences(0)
	return subsequences
}

/**
跟子集II的最大区别：
- 不能先做排序；
- 收集结果时需要限制一下长度。
*/

/**
两种去重逻辑：
- 记录树层——每层重新定义去重记录，无需随着递归一起回溯；
- 记录树枝——全局共用一个去重记录，需要随着递归一起回溯。

由于候选集有重复元素且结果不能重复，used记录的必须是某个元素是否已被选取，即used[nums[i]]，而不是某下标的元素是否已被选取，即used[i]。
*/
