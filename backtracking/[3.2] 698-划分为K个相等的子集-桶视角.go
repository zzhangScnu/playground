package backtracking

// 给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
// 示例 1：
//
// 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// 输出： True
// 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
// 示例 2:
//
// 输入: nums = [1,2,3,4], k = 3
// 输出: false
//
// 提示：
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
// 每个元素的频率在 [1,4] 范围内
func canPartitionKSubsetsII(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	var used int
	memo := make(map[int]bool)
	var traverse func(nums []int, start int, remainBucketNum int, remainTarget int) bool
	traverse = func(nums []int, start int, remainBucketNum int, remainTarget int) bool {
		if remainBucketNum == 0 {
			return true
		}
		if remainTarget < 0 {
			return false
		}
		if remainTarget == 0 {
			flag := traverse(nums, 0, remainBucketNum-1, target)
			memo[used] = flag
			return flag
		}
		if flag, ok := memo[used]; ok {
			return flag
		}
		for i := start; i < len(nums); i++ {
			if (used>>i)&1 == 1 {
				continue
			}
			used = used | 1<<i
			if traverse(nums, i+1, remainBucketNum, remainTarget-nums[i]) {
				return true
			}
			used = used ^ 1<<i
		}
		return false
	}
	return traverse(nums, 0, k, target)
}
