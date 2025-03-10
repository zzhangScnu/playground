package dynamicprogramming

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
// 序列。
//
// 示例 1：
//
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
// 示例 2：
//
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
// 示例 3：
//
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
// 进阶：
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
func lengthOfLISII(nums []int) int {
	var piles []int
	for _, num := range nums {
		index := searchFirstGE(piles, num)
		if index >= len(piles) {
			piles = append(piles, num)
		} else {
			piles[index] = num
		}
	}
	return len(piles)
}

func searchFirstGE(piles []int, num int) int {
	low, high := 0, len(piles)-1
	for low <= high {
		mid := low + (high-low)/2
		if piles[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

/**
贪心+二分查找
思路(从patient sort而来)：
维护一个牌堆，每个牌堆都是有序的，最上面的最小，往下相等或递增；
遍历候选集nums(1, 3, 5, 2, 6)，在牌堆中进行二分查找，每次都去找大于等于nums的第一个数，该位置的索引为index。
此时有两种情况：
情况一：
num = 2
piles = (1, 3, 5)
index = 1
所以用2代替3，piles = (1, 2, 5)
这里体现了贪心的思想，对于递增序列来说，(1, 2)优于(1, 3)，因为2相比于3更小，
为后续的数留有更多余地，有更多可能性组成更长的序列。

情况二：
num = 6
piles = (1, 2, 5)
index = 3
没有一个数字比6更大、能被6代替，所以扩展递增序列长度，
piles = (1, 2, 5, 6)

由此可见，piles维护的并不是实际的递增序列本身，
如对于6来说，长度为4时实际上的递增序列是(*, *, *, 6)，正确的结果是(1, 3, 5, 6)。

nums处理完成后，牌堆的数量就是递增序列的长度。直接返回即可。
同时，若将每个牌堆进行归并排序，则可得到一个最终有序的序列。
*/
