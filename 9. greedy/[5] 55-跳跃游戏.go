package greedy

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//
// 示例 2：
//
// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
// 提示：
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 10⁵
func canJump(nums []int) bool {
	var covered int
	for i := 0; i <= covered; i++ {
		covered = max(covered, i+nums[i])
		if covered >= len(nums) {
			return true
		}
	}
	return false
}

/**
局部最优：每次覆盖范围尽可能大；
全局最优：到达最后一个下标。
而不需要思考具体跳跃的路径，和跳跃几次才能到达。

通过max(covered, i+nums[i])，
在当前范围中，每一个被覆盖的下标上，计算下一次覆盖的最大范围。
covered就是当前所能覆盖到的最大下标。
*/
