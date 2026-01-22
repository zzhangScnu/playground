package slidingwindow

import "math"

// 给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。
//
// 找出满足下述条件的下标对 (i, j)：
//
// i != j,
// abs(i - j) <= indexDiff
// abs(nums[i] - nums[j]) <= valueDiff
//
// 如果存在，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
// 输出：true
// 解释：可以找出 (i, j) = (0, 3) 。
// 满足下述 3 个条件：
// i != j --> 0 != 3
// abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
// abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0
//
// 示例 2：
//
// 输入：nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
// 输出：false
// 解释：尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。
//
// 提示：
//
// 2 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 1 <= indexDiff <= nums.length
// 0 <= valueDiff <= 10⁹
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	bucketSize := valueDiff + 1
	bucket := make(map[int]int)
	var bucketId int
	for i, num := range nums {
		bucketId = calculateBucketId(bucketSize, num)
		if val, ok := bucket[bucketId]; ok && abs(num, val) <= valueDiff {
			return true
		}
		if val, ok := bucket[bucketId-1]; ok && abs(num, val) <= valueDiff {
			return true
		}
		if val, ok := bucket[bucketId+1]; ok && abs(num, val) <= valueDiff {
			return true
		}
		bucket[bucketId] = num
		if i >= indexDiff {
			bucketId = calculateBucketId(bucketSize, nums[i-indexDiff])
			delete(bucket, bucketId)
		}
	}
	return false
}

func calculateBucketId(bucketSize, num int) int {
	return int(math.Floor(float64(num) / float64(bucketSize)))
}

func abs(num1, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		return -diff
	}
	return diff
}

/**
题解：
indexDiff：滑动窗口控制
valueDiff：桶控制

桶详解
什么作用
划分数据范围，将元素划分到不同的桶中。在同一个桶中的元素，它们的差值都在固定范围内。桶大小为2，则每个桶的数据范围为[-1,0]和[0,1]，即桶内的元素最大差值为1。
相当于使用空间换时间，原本暴力解法一一对比是 O(n^2) 的时间复杂度，可以降为 O(n)。

如何划分
无论是正数还是负数，都需要【向下取整】。
对于 Go 来说，普通除法是向零取整的，因此需要用到 math.Floor 函数。

如何使用
- 如果两个数在同一个大小为 bucketSize 的桶中，说明这两个数的最大差值为 bucketSize - 1；
- 如果两个数不在同一个桶中，则有可能被划分到相邻的桶中了，此时需要分别判断左边和右边的桶中是否有差值满足要求的元素。

小优化
abs：原本用的 math.Abs 函数，但需要在整形和浮点数之间相互转换，效率低下。可以用判断正负 + 按需取反的方式平替。
*/
