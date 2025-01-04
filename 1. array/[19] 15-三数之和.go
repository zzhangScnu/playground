package array

import "slices"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
// k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 示例 1：
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 示例 2：
//
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0 。
//
// 示例 3：
//
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0 。
//
// 提示：
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var res [][]int
	for i := 0; i < len(nums); {
		num := nums[i]
		for _, twoSumRes := range twoSumOfDuplicatedNumbers(nums, i+1, 0-num) {
			res = append(res, append(twoSumRes, num))
		}
		for i < len(nums) && nums[i] == num {
			i++
		}
	}
	return res
}

func twoSumOfDuplicatedNumbers(nums []int, startIdx int, target int) [][]int {
	var res [][]int
	low, high := startIdx, len(nums)-1
	for low < high {
		lowVal, highVal := nums[low], nums[high]
		if lowVal+highVal == target {
			res = append(res, []int{lowVal, highVal})
			for low < high && nums[low] == lowVal {
				low++
			}
			for low < high && nums[high] == highVal {
				high--
			}
		} else if lowVal+highVal > target {
			high--
		} else {
			low++
		}
	}
	return res
}

/**
思路：
1. 创造数组非递减有序条件；
2. 固定第一个值，剩余的两个值用2Sum思路解决。

注意点1：For的写法
for i, num := range nums {
	for _, twoSumRes := range twoSumOfDuplicatedNumbers(nums, i+1, len(nums)-1, 0-num) {
		res = append(res, append(twoSumRes, num))
	}
	for i < len(nums) && nums[i] == num {
		i++
	}
}
这种写法是不work的，for range时是将i和num赋值给固定地址的局部变量，且在循环开始时，遍历的数据就已经固定了。

注意点2：非重复组合的控制
1. 对于剩余的两个值，可以在2Sum方法里控制；
2. 与之相对，第一个值需要在外层方法额外处理。

注意点3：重复元素的跳过
for i < len(nums) && nums[i] == num
for low < high && nums[low] == lowVal
1. 用局部变量，固定基准值；
2. 从当前值本身开始对比，这样即使后面没有重复元素，也能使得当前指针往前进一步，从下个位置重新开始；
3. 控制边界，避免越界。
*/
