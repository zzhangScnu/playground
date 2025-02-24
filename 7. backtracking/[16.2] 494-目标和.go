package backtracking

import "strconv"

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
func findTargetSumWaysII(nums []int, target int) int {
	var cnt int
	var traverse func(nums []int, index, sum int)
	traverse = func(nums []int, index, sum int) {
		if index == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}
		traverse(nums, index+1, sum+nums[index])
		traverse(nums, index+1, sum-nums[index])
	}
	traverse(nums, 0, 0)
	return cnt
}

func findTargetSumWaysWithMemo(nums []int, target int) int {
	memo := make(map[string]int)
	var traverse func(nums []int, index, sum int) int
	traverse = func(nums []int, index, sum int) int {
		key := strconv.Itoa(index) + "/" + strconv.Itoa(sum)
		if val, ok := memo[key]; ok {
			return val
		}
		if index == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		res := traverse(nums, index+1, sum+nums[index]) + traverse(nums, index+1, sum-nums[index])
		memo[key] = res
		return res
	}
	return traverse(nums, 0, 0)
}

/**
这种解法跟16.1的解法不一样的是，真的在每个元素前面尝试加上"+"或"-"，再看表达式计算结果是否满足题目中给的target。
所以需要索引遍历到候选集末尾，再来收集结果。且收集完后，后面已经没有可用元素了，需要直接返回，而不是继续向下递归。

备忘录的解法，多数需要递归函数有返回值，通过查找备忘录来提前获取&返回此次计算结果。
*/
