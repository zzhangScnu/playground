package backtracking

import "slices"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
// 示例 1：
//
// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
// [1,2,1],
// [2,1,1]]
//
// 示例 2：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 提示：
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
func permuteUnique(nums []int) [][]int {
	var path []int
	var res [][]int
	used := make([]bool, len(nums))
	var doPermute func(used []bool)
	doPermute = func(used []bool) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		usedOfLevel := make(map[int]interface{})
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if _, ok := usedOfLevel[nums[i]]; ok {
				continue
			}
			usedOfLevel[nums[i]] = true
			used[i] = true
			path = append(path, nums[i])
			doPermute(used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	doPermute(used)
	return res
}

func permuteUniqueSimplified(nums []int) [][]int {
	slices.Sort(nums)
	var path []int
	var res [][]int
	used := make([]bool, len(nums))
	var doPermute func(used []bool)
	doPermute = func(used []bool) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			doPermute(used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	doPermute(used)
	return res
}

/**
做法1：
在全排列的基础上，增加一个记录树层重复的数组，控制横向选取不重复；

做法2：
used有两个作用：
1. 树层上重复(参考组合总和II)：!used[i-1]
2. 树枝上重复：used[i] == true
*/
