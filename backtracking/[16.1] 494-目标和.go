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
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
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
用等式处理一下要求的值，转换为组合总和题目：
+sum(A) - sum(B) = target
sum(A) = target + sum(B)
sum(A) + sum(A) = target + sum(B) + sum(A)
2 * sum(A) = target + sum(nums)
+sum(A) = (target + sum(nums)) / 2
即求解有多少个+使得总和为(target + sum(nums)) / 2。

有重复元素，但不能重复选同一个。
需要注意的是，一开始写的base case是：
if start == len(nums) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	return
}
是有问题的，因为索引不一定要遍历完候选集才能找到一组满足和为(sum + target) / 2的组合，后面有些元素也许不选也能满足条件。
而且收集元素后也不能return，因为后面还有可能性，需要继续向后递归。
*/

// todo：slices.Sort(nums)，为什么要排序？
