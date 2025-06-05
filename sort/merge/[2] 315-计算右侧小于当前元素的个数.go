package merge

// 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
// nums[i] 的元素的数量。
//
// 示例 1：
//
// 输入：nums = [5,2,6,1]
// 输出：[2,1,1,0]
// 解释：
// 5 的右侧有 2 个更小的元素 (2 和 1)
// 2 的右侧仅有 1 个更小的元素 (1)
// 6 的右侧有 1 个更小的元素 (1)
// 1 的右侧有 0 个更小的元素
//
// 示例 2：
//
// 输入：nums = [-1]
// 输出：[0]
//
// 示例 3：
//
// 输入：nums = [-1,-1]
// 输出：[0,0]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴

var tempInCountSmaller [][]int

var count []int

func countSmallerII(nums []int) []int {
	var arr [][]int
	for i, num := range nums {
		arr = append(arr, []int{i, num})
	}
	n := len(nums)
	tempInCountSmaller = make([][]int, n)
	count = make([]int, n)
	sortInCountSmaller(arr, 0, n-1)
	return count
}

func sortInCountSmaller(arr [][]int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)>>1
	sortInCountSmaller(arr, lo, mid)
	sortInCountSmaller(arr, mid+1, hi)
	mergeInCountSmaller(arr, lo, mid, hi)
}

func mergeInCountSmaller(arr [][]int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		tempInCountSmaller[i] = arr[i]
	}
	i, j := lo, mid+1
	for cur := lo; cur <= hi; cur++ {
		if i == mid+1 {
			arr[cur] = tempInCountSmaller[j]
			j++
		} else if j == hi+1 {
			arr[cur] = tempInCountSmaller[i]
			count[tempInCountSmaller[i][0]] += j - mid - 1
			i++
		} else if tempInCountSmaller[i][1] <= tempInCountSmaller[j][1] {
			arr[cur] = tempInCountSmaller[i]
			count[tempInCountSmaller[i][0]] += j - mid - 1
			i++
		} else {
			arr[cur] = tempInCountSmaller[j]
			j++
		}
	}
}

/**
思路：
在归并排序中，合并两个已排序的子数组arr1和arr2到nums时，
若arr1[i] <= arr[j]，需将arr1[i]赋予nums[cur]，即nums[cur] = arr1[i]。
如：
nums = 1, 2, 3, 4, cur
arr1 = 1, 3, 5
arr2 = 2, 4, 6, 7
实际arr1和arr2是逻辑上的两个子数组，物理上同属于一个数组arr，由几个分隔索引切割为arr[lo, mid]和arr[mid + 1, hi]。
而处理arr1[i]时，可知arr2中前一部分的元素都是比arr1[i]小的。这些元素在原数组nums中位于arr1[i]的右侧（因为分治处理时分配到了右子数组中），故它们的个数 == arr1[i]右侧小于当前元素的个数。
这一部分元素的索引区间为[mid + 1, j)。

注意，arr1[i] <= arr[j]需包括 == 的情况，否则对于相等的情况，计算个数的操作会被跳过。

因为在归并处理过程中，索引变动频繁、与原数组nums无法直接关联，故需要自行维护nums中元素的[索引, 元素值]，基于其进行排序和处理。
*/
