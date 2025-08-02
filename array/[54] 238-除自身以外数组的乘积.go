package array

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
//
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在 32 位 整数范围内。
//
// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4]
// 输出: [24,12,8,6]
//
// 示例 2:
//
// 输入: nums = [-1,1,0,-3,3]
// 输出: [0,0,9,0,0]
//
// 提示：
//
// 2 <= nums.length <= 10⁵
// -30 <= nums[i] <= 30
// 输入 保证 数组 answer[i] 在 32 位 整数范围内
//
// 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0] = 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = prefix[i] * suffix[i]
	}
	return res
}

// todo：
// 在计算前缀乘积（`prefix`）和后缀乘积（`suffix`）时，之所以不需要包含当前位置的元素，是由「除自身以外数组的乘积」这一需求决定的。我们以 `[1,2,3,4]` 为例，具体解释：
//
//
//### 核心逻辑：结果 = 左侧乘积 × 右侧乘积
//对于数组中第 `i` 个元素 `nums[i]`，「除自身以外的乘积」= **左侧所有元素的乘积** × **右侧所有元素的乘积**。
//因此，`prefix` 和 `suffix` 的设计目标是：
//- `prefix[i]`：仅包含 `nums[0]` 到 `nums[i-1]` 的乘积（**不包含 `nums[i]`**）
//- `suffix[i]`：仅包含 `nums[i+1]` 到 `nums[n-1]` 的乘积（**不包含 `nums[i]`**）
//
//
//### 为什么「最后一个元素」不参与自身的计算？
//以最后一个元素 `nums[3] = 4` 为例：
//- 它的「除自身以外的乘积」是 `1×2×3 = 6`（即左侧所有元素的乘积）。
//- 此时，`prefix[3]` 已经包含了 `nums[0]×nums[1]×nums[2] = 6`（恰好是左侧所有元素的乘积）。
//- 而它的右侧没有元素（`i+1 = 4` 超出数组范围），因此 `suffix[3]` 设为 `1`（乘法的「单位元」，不影响结果）。
//- 最终 `res[3] = prefix[3] × suffix[3] = 6 × 1 = 6`，正确排除了自身 `4`。
//
//
//### 扩展到所有元素
//对于任意位置 `i`：
//- 如果包含 `nums[i]` 到 `prefix` 或 `suffix` 中，结果会混入自身的值，违反「除自身以外」的要求。
//- 前缀和后缀的「不包含当前元素」的设计，正是为了精准计算「左侧×右侧」的乘积。
//
//
//简言之，不是「不需要计算最后一个 `nums`」，而是**所有位置的 `nums[i]` 都不参与自身的前缀/后缀计算**，这样才能确保结果是「除自身以外的乘积」。

func productExceptSelfII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}
	return res
}
