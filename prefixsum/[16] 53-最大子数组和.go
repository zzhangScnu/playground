package prefixsum

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。
//
// 示例 1：
//
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组[4,-1,2,1] 的和最大，为6 。
//
// 示例 2：
//
// 输入：nums = [1]
// 输出：1
//
// 示例 3：
//
// 输入：nums = [5,4,-1,7,8]
// 输出：23
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
// 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
func maxSubArray(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	minPrefixSum, res := prefixSum[0], nums[0]
	for i := 1; i <= n; i++ {
		res = max(res, prefixSum[i]-minPrefixSum)
		minPrefixSum = min(minPrefixSum, prefixSum[i])
	}
	return res
}

// todo：计算顺序的关键在于保证子数组长度≥1：
//
//1. 时间维度 ：
//
//   - prefixSum[i] - minPrefixSum 表示从 j+1 到 i 的子数组和
//   - 必须保证 j < i → minPrefixSum 必须来自 prefixSum[0..i-1]
