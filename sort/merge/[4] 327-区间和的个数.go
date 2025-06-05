package merge

// 给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和
// upper）之内的 区间和的个数 。
//
// 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
//
// 示例 1：
//
// 输入：nums = [-2,5,-1], lower = -2, upper = 2
// 输出：3
// 解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
//
// 示例 2：
//
// 输入：nums = [0], lower = 0, upper = 0
// 输出：1
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// -10⁵ <= lower <= upper <= 10⁵
// 题目数据保证答案是一个 32 位 的整数

var rangeSumCount int

var tempInCountRangeSum []int

func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int, len(nums)+1)
	rangeSumCount, tempInCountRangeSum = 0, make([]int, len(preSum))
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	sortInCountRangeSum(preSum, lower, upper, 0, len(preSum)-1)
	return rangeSumCount
}

func sortInCountRangeSum(preSum []int, lower, upper int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)>>1
	sortInCountRangeSum(preSum, lower, upper, low, mid)
	sortInCountRangeSum(preSum, lower, upper, mid+1, high)
	mergeInCountRangeSum(preSum, lower, upper, low, mid, high)
}

func mergeInCountRangeSum(preSum []int, lower, upper int, low, mid, high int) {
	left, right := mid+1, mid+1
	for cur := low; cur <= mid; cur++ {
		for left <= high && preSum[left]-preSum[cur] < lower {
			left++
		}
		for right <= high && preSum[right]-preSum[cur] <= upper {
			right++
		}
		rangeSumCount += right - left
	}
	for i := low; i <= high; i++ {
		tempInCountRangeSum[i] = preSum[i]
	}
	i, j := low, mid+1
	for cur := low; cur <= high; cur++ {
		if i == mid+1 {
			preSum[cur] = tempInCountRangeSum[j]
			j++
		} else if j == high+1 {
			preSum[cur] = tempInCountRangeSum[i]
			i++
		} else if tempInCountRangeSum[i] <= tempInCountRangeSum[j] {
			preSum[cur] = tempInCountRangeSum[i]
			i++
		} else {
			preSum[cur] = tempInCountRangeSum[j]
			j++
		}
	}
}

/**
思路：归并排序
在归并排序的两个阶段：分治 - 合并之间，
可以夹带私货，利用2个已经各自有序的左数组nums[low, mid]和右数组nums[mid + 1, high]，实现目标。

对于本题，可以先基于原数组nums计算出前缀和数组preSum，再对preSum进行归并排序。
对于preSum的右数组来说，物理位置整体位于左数组的右侧。即左数组的每一个元素i和右数组的每一个元素j构成的preSum[j] - preSum[i]，都是原数组中的合法区间和。

则对于左数组preSum[low, mid]的每一个cur，寻找右数组preSum[mid + 1, high]中的j，使得lower <= preSum[j] - preSum[cur] <= upper。
如果此时用两层for循环依次检查，则时间复杂度O(n^2)。
但因为子数组的有序性，可以利用类滑动窗口思想，
对于preSum[low, mid]的每一个cur，在preSum[mid + 1, high]中找到一个尽可能大的窗口[left, right]，使得窗口内的元素与cur的差值落在[lower, upper]区间内：
- 如果preSum[left] - preSum[cur] < lower，则增大left，使差值落入左边界lower的右侧，即满足preSum[left] - preSum[cur] >= lower；
- 如果preSum[right] - preSum[cur] > upper，则缩小right，使差值落入右边界upper的左侧，即满足preSum[right] - preSum[cur] <= upper。
对下一个元素cur，因为子数组有序，preSum[cur]增大。
所以preSum[left]也需要增大，即left不动或向右查找，才能满足preSum[left] - preSum[cur] >= lower；
而preSum[mid + 1, right]肯定都满足preSum[right] - preSum[cur] <= upper，此时不动或向右推进right，寻找更多可能性。
即当cur右移时，相应的窗口也需在有序数组中整体右移。
*/
