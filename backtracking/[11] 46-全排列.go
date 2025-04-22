package backtracking

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 示例 2：
//
// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
//
// 示例 3：
//
// 输入：nums = [1]
// 输出：[[1]]
//
// 提示：
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
func permute(nums []int) [][]int {
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
- 叶子节点收集结果；
- 由于(1, 2)和(2, 1)是两个不同的结果，每层都需要从0开始遍历取元素，但需要记录曾经被选取过的元素，避免重复；
- used的作用是提供本层剩余可选集：如候选集[1, 2, 3]：
  选择1后，used = [1, 0, 0]；
  递归进入下一层后，由used可知剩余可选元素为[2, 3]。
*/
