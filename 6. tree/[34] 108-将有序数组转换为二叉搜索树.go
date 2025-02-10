package tree

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 平衡 二叉搜索树。
//
// 示例 1：
//
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
//
// 示例 2：
//
// 输入：nums = [1,3]
// 输出：[3,1]
// 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
//
// 提示：
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 按 严格递增 顺序排列
func sortedArrayToBST(nums []int) *TreeNode {
	var traverse func(nums []int, start, end int) *TreeNode
	traverse = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		val := nums[mid]
		return &TreeNode{
			Val:   val,
			Left:  traverse(nums, start, mid-1),
			Right: traverse(nums, mid+1, end),
		}
	}
	return traverse(nums, 0, len(nums)-1)
}
