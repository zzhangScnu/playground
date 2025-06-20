package backtracking

import "slices"

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
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
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
	slices.Sort(nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	bucketSum := make([]int, k)
	var traverse func(nums []int, numIndex int) bool
	traverse = func(nums []int, numIndex int) bool {
		if numIndex == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			if bucketSum[i]+nums[numIndex] > target {
				continue
			}
			if i > 0 && bucketSum[i] == bucketSum[i-1] {
				continue
			}
			bucketSum[i] += nums[numIndex]
			if traverse(nums, numIndex+1) {
				return true
			}
			bucketSum[i] -= nums[numIndex]
			if bucketSum[i] == 0 {
				break // todo:???
			}
		}
		return false
	}
	return traverse(nums, 0)
}

/**
排列组合问题可以抽象为球盒模型，将N个不同序号的球放入K个不同序号的桶中，其中N >= K。
其中有两种视角：
1. 站在数字视角，每个数字都要寻找一个桶放入；
   即第一个数字可以遍历K个桶，选择是否进入桶中。剩下N-1个球可以在K个桶中继续做选择。
2. 站在桶视角，每个桶都需要选择至少一个球。
   即第一个桶可以遍历N个球，选择是否将当前数字装进桶中。剩下K-1个桶可以在N-1个球中继续做选择。


思路：数字视角
设有K个桶，每个桶的容量为target。
对于数组中的每一个数字，可以有K个选择，放入[0, K-1]区间的桶中。

仅使用for循环时：
// 穷举 nums 中的每个数字
for (int i = 0; i < len(nums); i++) {
	// 穷举每个桶
	for (int j = 0; j < k; j++) {
		// nums[i] 选择是否要进⼊第 j 个桶
		// ...
	}
}

改为递归形式：
// 穷举 nums 中的每个数字
func backtrack(nums []int, index int, bucket []int) {
    // base case
    if index == len(nums) {
        return
    }
    // 穷举每个桶
    for i := 0; i < len(bucket); i++ {
        // 选择装进第i个桶
        bucket[i] += nums[index]
        // 递归穷举下一个数字的选择
        backtrack(nums, index+1, bucket)
        // 撤销选择
        bucket[i] -= nums[index]
    }
}

if traverse(nums, numIndex+1) {
	return true
}
分支搜索，找到回溯树中任一符合条件路径即终止搜索，向上层层返回。

优化：
1. 将数组倒序排序，快速将不合法情况暴露出来，节省后续无用尝试，即尽可能多地命中下面的剪枝条件；
2. 由于数组有序，如果有重复元素，则回溯树的相邻相等树枝生长出来的子树会相等。
   此时将后续子树剪枝，只遍历第一棵子树即可。

时间复杂度：
设 N = len(nums)，K = 桶的数量
每个数字有K个选择，则N个数字：
O(K^N)

跟桶视角相比可知，要尽量收缩可选择范围，因为每多一个选项都是指数级别增长；
相比之下只增加候选集个数(桶的数量)，即多做几次选择，只是常数级别增长。
*/
