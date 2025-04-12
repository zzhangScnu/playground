package array

// 给定一个非负整数数组 nums 和一个整数 k ，你需要将这个数组分成 k 个非空的连续子数组，使得这 k 个子数组各自和的最大值 最小。
//
// 返回分割后最小的和的最大值。
//
// 子数组 是数组中连续的部份。
//
// 示例 1：
//
// 输入：nums = [7,2,5,10,8], k = 2
// 输出：18
// 解释：
// 一共有四种方法将 nums 分割为 2 个子数组。
// 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
// 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
//
// 示例 2：
//
// 输入：nums = [1,2,3,4,5], k = 2
// 输出：9
//
// 示例 3：
//
// 输入：nums = [1,4,4], k = 3
// 输出：4
//
// 提示：
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 10⁶
// 1 <= k <= min(50, nums.length)
func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, num := range nums {
		left = max(left, num)
		right += num
	}
	for left <= right {
		mid := left + (right-left)>>1
		if canSplit(nums, mid, k) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canSplit(nums []int, targetSum int, k int) bool {
	sum, count := 0, 1
	for _, num := range nums {
		if sum+num <= targetSum {
			sum += num
		} else {
			count++
			sum = num
		}
	}
	return count <= k
}

/**
题目：给定一个非负整数数组nums和一个整数k，将这个数组分成k个非空的连续子数组，使得这k个子数组各自和的最大值最小。
		  求分割后最小的和的最大值。

分割k次后，各连续子数组总和sum的最大值，若要达到所有情况中的最小，应满足各个sum较为相近，分布均匀。
题目转化为：
是否存在一个最小值targetSum，使得各子数组总和sum <= targetSum，且分割次数 <= k。
因为子数组均满足总和sum <= targetSum，该targetSum即为题目所求的解。
注意，题目中声明"将数组分成k个非空的连续子数组"，即分割次数 == 子数组数量，
即最小分割次数 == 1。

为什么是 分割次数 <= k，而不是 == k？
1. 假如分割了k - 1次，则还可以将sum较大的子数组进一步分割，分割后各子数组的sum仍满足 <= targetSum的约束；
   此时可以进一步缩小targetSum，试图寻找最小值。
2. 如果限定分割k次，会导致场景遗漏：
   若nums的总和 <= targetSum，此时分割次数为1，很有可能小于题目中给的k。如果限定分割k次，则会拒绝这种场景。
   但实际上可以进一步将nums分割，并缩小targetSum，寻找可能的解。


三部曲：

1. 画出函数在二维坐标上的图像，明确 x、f(x)、target，并实现函数 f；
   - x：各子数组总和的最大值；
   - f(x)：是否存在分割次数 <= k，使得各子数组总和sum <= x；
   - target：是，即f(x) == true。

2. 明确 x 的取值范围，作为二分搜索的搜索区间，初始化left和right变量；
	- left：nums中的最大值。此时每个元素自身组成子数组，满足sum <= x；
	- right：nums的总和。此时nums自身组成子数组，满足sum <= x。

3. 根据题意明确使用搜索左侧 / 右侧的二分搜索算法，写出解法代码。
	在f(x) == true的约束下，对x进行二分搜索，寻找左边界，即targetSum的最小值。
*/
