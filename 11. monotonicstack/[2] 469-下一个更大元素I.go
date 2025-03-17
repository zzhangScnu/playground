package monotonicstack

// nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
//
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
//
// 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定
// nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
//
// 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
//
// 示例 1：
//
// 输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出：[-1,3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// - 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
// - 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
// - 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//
// 示例 2：
//
// 输入：nums1 = [2,4], nums2 = [1,2,3,4].
// 输出：[3,-1]
// 解释：nums1 中每个值的下一个更大元素如下所述：
// - 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
// - 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
//
// 提示：
//
// 1 <= nums1.length <= nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 10⁴
// nums1和nums2中所有整数 互不相同
// nums1 中的所有整数同样出现在 nums2 中
//
// 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	posMap := make(map[int]int)
	for index, num := range nums1 {
		posMap[num] = index
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = -1
	}
	st := make([]int, len(nums2))
	st[0] = 0
	for i := 1; i < len(nums2); i++ {
		if nums2[i] <= nums2[st[len(st)-1]] {
			st = append(st, i)
		} else {
			for len(st) > 0 && nums2[i] > nums2[st[len(st)-1]] {
				if posInNums1, ok := posMap[nums2[st[len(st)-1]]]; ok {
					res[posInNums1] = i
				}
				st = st[:len(nums2)-1]
			}
			st = append(st, i)
		}
	}
	return res
}

/**
思路：
遍历nums2 -> 通过单调栈，找每个num下一个更大元素 -> 将num映射回nums1 -> 记录结果

因为是要找nums2中每个元素的下一个最大元素 -> 所以要遍历nums2 -> 所以栈的大小应初始化为nums2的大小
因为当右侧没有比当前元素更大的元素时，应赋值-1 -> 所以结果集初始化为-1
因为nums1是nums2的子集 -> 所以不是每一个nums2中的元素在nums1中都能找到 -> 所以在从nums2映射到nums1时，要先判断存在与否
因为当前元素跟栈口元素比较完后，还要跟nums2中的后来者继续比较 -> 所以记得要入栈
*/
