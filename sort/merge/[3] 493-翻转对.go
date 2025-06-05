package merge

// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
//
// 你需要返回给定数组中的重要翻转对的数量。
//
// 示例 1:
//
// 输入: [1,3,2,3,1]
// 输出: 2
//
// 示例 2:
//
// 输入: [2,4,3,5,1]
// 输出: 3
//
// 注意:
//
// 给定数组的长度不会超过50000。
// 输入数组中的所有数字都在32位整数的表示范围内。
var reversePairCount int

var tempInReversePairs []int

func reversePairs(nums []int) int {
	reversePairCount = 0
	tempInReversePairs = make([]int, len(nums))
	sortInReversePairs(nums, 0, len(nums)-1)
	return reversePairCount
}

func sortInReversePairs(nums []int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)>>1
	sortInReversePairs(nums, low, mid)
	sortInReversePairs(nums, mid+1, high)
	mergeInReversePairs(nums, low, mid, high)
}

func mergeInReversePairs(nums []int, low, mid, high int) {
	j := mid + 1
	for i := low; i <= mid; i++ {
		for ; j <= high && nums[i] > 2*nums[j]; j++ {
		}
		reversePairCount += j - mid - 1
	}
	for i := low; i <= high; i++ {
		tempInReversePairs[i] = nums[i]
	}
	i, j := low, mid+1
	for cur := low; cur <= high; cur++ {
		if i == mid+1 {
			nums[cur] = tempInReversePairs[j]
			j++
		} else if j == high+1 {
			nums[cur] = tempInReversePairs[i]
			i++
		} else if tempInReversePairs[i] <= tempInReversePairs[j] {
			nums[cur] = tempInReversePairs[i]
			i++
		} else {
			nums[cur] = tempInReversePairs[j]
			j++
		}
	}
}

/**
思路一：暴力
对于每一个i < j，判断是否满足nums[i] > 2*nums[j]。
时间复杂度O(n^2)。

思路二：归并排序
在归并排序的两个阶段：分治 - 合并之间，
可以夹带私货，利用2个已经各自有序的左数组nums[low, mid]和右数组nums[mid + 1, high]，实现目标。

对于本题来说，右数组中的每一个nums[j]都比左数组中的任意nums[i]更靠右、索引更大，即满足i < j，
在此基础上检查是否满足nums[i] > 2*nums[j]约束即可。

直观的思路是用嵌套的for循环进行一一对比，但容易超时。

此时可以利用元素有序性：对于一对i和j，若nums[i] > 2*nums[j]，则nums[i] > 2*nums[mid + 1 ... j]，即满足约束的元素个数为j - (mid + 1)
则维护一个count，固定每一个i，寻找满足条件的j的边界，当不再满足时意味着摸到头了，进行count += j - mid - 1。
这里有个小技巧，无需在切到下一个i后，每次都从mid + 1开始判断约束成立。
因为对于上一轮的i和j，已经满足nums[i] > 2*nums[j]；这一轮i++后，nums[i + 1] > 2*nums[j]自然成立，所以j应该也要基于上一轮的j往后寻找符合条件的位置。
时间复杂度O(NlogN)。
*/
