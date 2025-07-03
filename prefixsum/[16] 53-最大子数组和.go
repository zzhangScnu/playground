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

/**
前缀和数组做法

思路：
先将每个位置组成子数组的和计算出来，即prefixSum[i]表示索引在[1, i]范围内的元素之和。
prefixSum从1开始，prefixSum[0] = 0为base case，为了计算前缀和数组时更方便，无需额外考虑数组越界的情况。
从索引1开始遍历前缀和数组，一边维护最小的前缀和，一边与当前位置的前缀和相减。
假设最小前缀和为prefixSum[i]，当前指针指向prefixSum[j]，则prefixSum[j]-prefixSum[i]表示索引在[i+1,...j]范围内的元素之和。
对这些差值取最大值，即为最大子数组和。

注意：
需先计算res，再计算minPrefixSum，因为需保证子数组长度 >= 1。
prefixSum[i] - minPrefixSum 表示索引范围为 [j, i] 的子数组和，必须保证 j < i。所以minPrefixSum 必须来自 prefixSum[0..i-1]，即在每轮计算完res后再进行更新。
否则如果当前位置的前缀和即为最小，在先计算minPrefixSum再计算res的情况下，计算出来的res为0，无法正确更新。

同时，因为数组中有负数，不能用滑动窗口解法；
且minPrefixSum应初始化为prefixSum[0]即0；
res应初始化为nums[0]，因子数组最少长度为1，即元素本身。
*/
