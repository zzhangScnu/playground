package dynamicprogramming

// 给定一个未排序的整数数组
// nums ， 返回最长递增子序列的个数 。
//
// 注意 这个数列必须是 严格 递增的。
//
// 示例 1:
//
// 输入: [1,3,5,4,7]
// 输出: 2
// 解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
//
// 示例 2:
//
// 输入: [2,2,2,2,2]
// 输出: 5
// 解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
//
// 提示:
//
// 1 <= nums.length <= 2000
// -10⁶ <= nums[i] <= 10⁶
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	count, length := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		count[i], length[i] = 1, 1
	}
	var maxLen int
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if length[j] == length[i] {
				count[i] += count[j]
			}
			if length[j]+1 == length[i] {
				count[i] = count[j]
				length[i] = length[j] + 1
				maxLen = max(maxLen, length[i])
			}
		}
	}
	var res int
	for i := 0; i < n; i++ {
		if length[i] == maxLen {
			res += count[i]
		}
	}
	return res
}

/**
DP数组及下标含义：
- i：当前游标指向位置；
- count[i]：nums[0 ... i]间最长递增子序列的个数。
- len[i]：nums[0 ... i]间最长递增子序列的长度。

递推公式：
对于case [1, 3, 5, 4, 7]，
for i -> [0, n)
for j -> [0, i)
固定i作为右边界，对每一个小于i的j进行判断：
若 nums[j] < nums[i]，则表示nums[j ... i]单调递增，
	- 若len[j] == len[i]，则表示找到了另一组长度相同的递增子序列。
	  形如[1, 3, 5] 和[1, 3, 4]。此时j = 2，i = 3，len[2] == len[3] == 3，count[2] = 1，
	  故count[i] += count[j]，len[i] = len[j]
	- 若len[j] + 1 == len[i]，则表示找到了另一组更长的递增子序列。
	  形如[1, 3, 5] 和[1, 3, 5, 7]。此时j = 2，i = 4，len[2] == 3，len[4] == 4，count[2] = 1，
 	  故count[i] = count[j]，len[i] = len[j] + 1

初始化：
count中每个单元格的值均为1，表示自身组成一个递增子序列。
len中每个单元格的值均为1，表示自身组成递增子序列的长度为1。

遍历方向：
先遍历i，再遍历j。固定i的情况下，对[0 ... i]中的每一个子序列[0 ... j]，判断和nums[i]之间是否构成递增关系。
从左到右。
*/
