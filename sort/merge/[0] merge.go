package merge

import "math"

func mergeSort(nums []int) []int {
	return doMergeSort(nums, 0, len(nums)-1)
}

func doMergeSort(nums []int, beginIdx, endIdx int) []int {
	if beginIdx == endIdx {
		return []int{nums[beginIdx]}
	}
	mid := beginIdx + (endIdx-beginIdx)/2
	nums1 := doMergeSort(nums, beginIdx, mid)
	nums2 := doMergeSort(nums, mid+1, endIdx)
	return merge(nums1, nums2)
}

func merge(nums1, nums2 []int) []int {
	INF := math.MaxInt64
	nums1, nums2 = append(nums1, INF), append(nums2, INF)
	var i, j int
	var res []int
	for nums1[i] != INF || nums2[j] != INF {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	return res
}

/**
归并排序：
- 思路：分而治之
   将问题以同样形式分解为子问题，在到达base case(触底)时，开始回溯(反弹)，逐一解决拆分后的子问题。
- 实现：递归
	使用游标来控制分解&回溯的子问题边界。
- 本质：二叉树的后序遍历。
	二叉树问题解法分为1. 遍历，2. 分解，归并排序利用的是分解问题的思路。
  - 递归算法本质在遍历一棵递归树，在节点的前/中/后序位置上执行代码。递归算法的核心就是明确每个节点上的单层逻辑；
  - 归并排序的过程可抽象为二叉树，每个节点的值为nums[beginIdx...endIdx]，则叶子节点的值为单个元素；
  - merge操作会在【每个节点】的后序遍历位置执行，合并两个子节点上的子数组。

	```
	func doMergeSort(nums []int, beginIdx, endIdx int) []int {
		if beginIdx == endIdx {
			return []int{nums[beginIdx]}
		}
		mid := beginIdx + (endIdx-beginIdx)/2
		nums1 := doMergeSort(nums, beginIdx, mid)
		nums2 := doMergeSort(nums, mid+1, endIdx)
		return merge(nums1, nums2)
	}

	和二叉树的后序遍历实现很像：

	func traverse(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		traverse(root.Right)
		// 后序位置
		// do something
	}
	```

哨兵解决边界问题：
普通实现方式，需要在较短数组处理完成后，将较长数组剩余元素逐一拷贝到结果集，需要处理数组越界情况；
但通过在原始数组末尾增加哨兵的形式，可以简化处理逻辑。

需要O(n)的额外空间复杂度，在每次合并过程中作为临时结果集。

平均时间复杂度为O(NlogN)。
- 树的高度：logN
- merge函数执行次数 == 二叉树节点个数 ->
  merge函数执行复杂度 == 每个节点代表的子数组的长度 ->
  每层总时间复杂度 == 二叉树本层节点个数 * 每个节点代表的子数组的长度 == 本层元素个数 ->
  每层元素个数 == 原数组长度n
- 故总时间复杂度 == O(N * logN)

归并排序是稳定排序，意味着原数组中相同元素的相对顺序在排序后仍然保持不变。
因为在merge方法中，左右两个已排序的数组仍保持在原本的数组中的相对顺序。
对相同元素合并时，按顺序取即可，可以人为控制稳定性。
*/
