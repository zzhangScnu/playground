package backtracking

import (
	"math"
	"slices"
)

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
	var path []int
	var res [][]int
	var traverse func(nums []int, target int, start int)
	traverse = func(nums []int, target int, start int) {
		if target < 0 {
			return
		}
		if start == len(nums) {
			if target == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			traverse(nums, target-nums[i], i+1)
			path = path[:len(path)-1]
		}
	}
	slices.Sort(nums)
	var sum int
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum < int(math.Abs(float64(target))) {
		return 0
	}
	target = (sum + target) / 2
	traverse(nums, target, 0)
	return len(res)
}

/**
有重复元素，但不能重复选同一个。
*/
