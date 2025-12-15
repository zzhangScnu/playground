package N数之和

import "slices"

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
// b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
// 你可以按 任意顺序 返回答案 。
//
// 示例 1：
//
// 输入：nums = [1,0,-1,0,-2,2], target = 0
// 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
// 示例 2：
//
// 输入：nums = [2,2,2,2,2], target = 8
// 输出：[[2,2,2,2]]
//
// 提示：
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	return nSum(nums, 4, 0, target)
}

func nSum(nums []int, n int, startIdx, target int) [][]int {
	if n == 2 {
		return twoSumOfDuplicatedNumbers(nums, startIdx, target)
	}
	var res [][]int
	for i := startIdx; i < len(nums)-n+1; {
		num := nums[i]
		for _, partRes := range nSum(nums, n-1, i+1, target-num) {
			res = append(res, append(partRes, num))
		}
		for i < len(nums)-n+1 && nums[i] == num {
			i++
		}
	}
	return res
}

/**
用递归来解决nSUM问题。
递归，不能代入去执行，而是从逻辑层面去理解：每一层需要做的事情 + 跳出的条件

1. 去重
- 用局部变量，固定基准值；
- 从当前值本身开始对比，这样即使后面没有重复元素，也能使得当前指针往前进一步，从下个位置重新开始（兼容当前位置在未重复情况下也能推进）。
2. 遍历范围：len(nums)-n+1
- n 数之和，意味着最后位置起码需要预留 n 个数字进行求和。
*/
