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
