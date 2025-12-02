package array

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
//
// 示例 2:
//
// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
// 进阶：
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
func rotate(nums []int, k int) {
	var times int
	n := len(nums)
	from, to := 0, 0
	for times < n {
		to = (from + k) % n
		nums[from], nums[to] = nums[to], nums[from]
		from = to
		times++
	}
}

func rotateII(nums []int, k int) {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[n-k:], nums[:n-k]...))
}

func rotateIII(nums []int, k int) {
	n := len(nums)
	if k%n == 0 {
		return
	}
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

/**
思路一：
将后半截数组拼接到前半截数组前面。
为什么用n-k？
本质上是数据移动方向。
截断&拼接的方式，是将后k个元素移动到左侧，将前n-k个元素移动到右侧。

思路二：
转换为多次反转数组。
注意base case：
若k%n == 0则不处理。这种情况即使操作轮转，结果跟原始数组也是一样的。


为什么方法一rotate不行？
有些case会导致在几个数字之间循环跳跃，无法处理其他元素。
*/
