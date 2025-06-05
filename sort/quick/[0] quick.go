package quick

func quickSort(nums []int) []int {
	doQuickSort(nums, 0, len(nums)-1)
	return nums
}

func doQuickSort(nums []int, beginIdx, endIdx int) {
	if beginIdx < 0 || endIdx > len(nums)-1 {
		return
	}
	if beginIdx >= endIdx {
		return
	}
	pivot := partition(nums, beginIdx, endIdx)
	doQuickSort(nums, beginIdx, pivot-1)
	doQuickSort(nums, pivot+1, endIdx)
}

func partition(nums []int, beginIdx, endIdx int) int {
	pivot := nums[endIdx]
	slow, fast := beginIdx, beginIdx
	for ; fast < endIdx; fast++ {
		if nums[fast] < pivot {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	nums[slow], nums[endIdx] = nums[endIdx], nums[slow]
	return slow
}

/**
原地排序算法，无需额外空间复杂度。
平均时间复杂度为O(NlogN)。

合并排序算法：从下往上，先分解，再合并
快速排序算法：从上往下，先分区，再分解
*/

/**
对pivot的移动，直接与slow交换，时间复杂度为O(1)，
而不需要将slow及之后的元素逐个往后搬移一位，导致上升到O(n)
*/

/**
快速排序是不稳定排序，因为在排序过程中，相同元素的相对位置可能会发生变化。
每轮选择一个基准元素pivot，然后将数组递归地划分为两个子数组，
左半子数组中的元素 <= 基准元素，右半子数组中的元素 >= 基准元素。
在这个过程中，相同元素可能会划分到不同的子数组中，也可以其一会被选中作为基准元素，无法人为控制相对顺序，也就导致了不稳定性。
*/

/**
本质是二叉树的前序遍历。
在前序位置选取出基准元素，
再递归处理左右数组。
当触底到达base case时，数组区间只有一个元素，直接返回。

而且这个二叉树是个二叉搜索树，严格遵循左子树（左数组）< 根节点（基准元素） < 右子树（右数组）的约束。
*/

/**
为了防止数组原本有序导致的时间复杂度由O(NlogN)退化为O(N^2)，可以通过两种方法规避：
1. 洗牌算法：排序前打乱数组元素的顺序；
2. 随机选择基准元素：每次递归时随机选择基准元素。
*/
