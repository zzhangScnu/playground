package N数之和

import (
	"math"
	"slices"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
// 示例 1：
//
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
//
// 示例 2：
//
// 输入：nums = [0,0,0], target = 1
// 输出：0
// 解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。
//
// 提示：
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	slices.Sort(nums)
	closestSubtraction, sum := math.MaxInt, 0
	for i := 0; i < len(nums)-2; i++ {
		sum = nums[i] + twoSumClosest(nums, i+1, target-nums[i])
		if math.Abs(float64(target-sum)) < math.Abs(float64(closestSubtraction)) {
			closestSubtraction = target - sum
		}
	}
	return target - closestSubtraction
}

func twoSumClosest(nums []int, start int, target int) int {
	closestSubtraction := math.MaxInt
	sum, low, high := 0, start, len(nums)-1
	for low < high {
		sum = nums[low] + nums[high]
		if math.Abs(float64(target-sum)) < math.Abs(float64(closestSubtraction)) {
			closestSubtraction = target - sum
		}
		if sum < target {
			low++
		} else {
			high--
		}
	}
	return target - closestSubtraction
}

/**
思路：
要找最接近的三数之和，可能是小于，也可能是大于。一开始的实现，就只考虑在小于target的情况下逼近target找最大值，是错误的。
本题应着眼于绝对值，寻找|target - 三数之和|最小的三元组。

使用双向指针的前提是数组有序，故先排序。
先固定第一个元素num，从num之后的位置开始下探寻找|target - 两数之和|最小的二元组。
因为用了一个变量closestSubtraction滚动维护最小的差值，closestSubtraction = target - sum，
所以最终返回的结果，移项可得sum = target - closestSubtraction，故返回target - closestSubtraction。

另外，因为最小绝对值可能出现在大于target，也可能出现在小于target时，
所以twoSumClosest中，最小绝对值的判断和更新，放在循环开始时，而不是依附于sum和target大小关系的不同分支中。

一开始的实现，会因为数值范围不匹配而溢出：
closestSubtraction := math.MaxInt
if int(math.Abs(float64(target)-float64(sum))) < int(math.Abs(float64(closestSubtraction)))
其中，int(math.Abs(float64(closestSubtraction)))，closestSubtraction -> float64 -> int，这个int取决于系统位数，不能保证是64位。
改写为直接使用浮点数比较，不会丢失精度或改变符号。
*/
