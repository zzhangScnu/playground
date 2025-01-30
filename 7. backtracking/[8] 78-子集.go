package backtracking

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
// 示例 1：
//
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
// nums 中的所有元素 互不相同【

var subset []int

var subsetsRes [][]int

func subsets(nums []int) [][]int {
	subset, subsetsRes = []int{}, [][]int{}
	doSubsets(nums, 0)
	return subsetsRes
}

func doSubsets(nums []int, beginIdx int) {
	//tmp := make([]int, len(subset))
	//copy(tmp, subset)
	//subsetsRes = append(subsetsRes, tmp)
	subsetsRes = append(subsetsRes, append([]int{}, subset...))
	//if beginIdx >= len(nums) {
	//	return
	//}
	for i := beginIdx; i < len(nums); i++ {
		subset = append(subset, nums[i])
		doSubsets(nums, i+1)
		subset = subset[:len(subset)-1]
	}
}
