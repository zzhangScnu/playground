package 滑动窗口

// 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于
// limit 。
//
// 如果不存在满足条件的子数组，则返回 0 。
//
// 示例 1：
//
// 输入：nums = [8,2,4,7], limit = 4
// 输出：2
// 解释：所有子数组如下：
// [8] 最大绝对差 |8-8| = 0 <= 4.
// [8,2] 最大绝对差 |8-2| = 6 > 4.
// [8,2,4] 最大绝对差 |8-2| = 6 > 4.
// [8,2,4,7] 最大绝对差 |8-2| = 6 > 4.
// [2] 最大绝对差 |2-2| = 0 <= 4.
// [2,4] 最大绝对差 |2-4| = 2 <= 4.
// [2,4,7] 最大绝对差 |2-7| = 5 > 4.
// [4] 最大绝对差 |4-4| = 0 <= 4.
// [4,7] 最大绝对差 |4-7| = 3 <= 4.
// [7] 最大绝对差 |7-7| = 0 <= 4.
// 因此，满足题意的最长子数组的长度为 2 。
//
// 示例 2：
//
// 输入：nums = [10,1,2,4,7,2], limit = 5
// 输出：4
// 解释：满足题意的最长子数组是 [2,4,7,2]，其最大绝对差 |2-7| = 5 <= 5 。
//
// 示例 3：
//
// 输入：nums = [4,2,2,2,4,4,2,2], limit = 0
// 输出：3
//
// 提示：
//
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^9
// 0 <= limit <= 10^9

type AscendingDeque []int

func (a *AscendingDeque) Len() int {
	return len(*a)
}

func (a *AscendingDeque) Push(num int) {
	for a.Len() > 0 && (*a)[a.Len()-1] > num {
		*a = (*a)[:a.Len()-1]
	}
	*a = append(*a, num)
}

func (a *AscendingDeque) Pop(num int) {
	if a.Len() > 0 && (*a)[0] == num {
		*a = (*a)[1:]
	}
}

func (a *AscendingDeque) Min() int {
	return (*a)[0]
}

type DescendingDeque []int

func (d *DescendingDeque) Len() int {
	return len(*d)
}

func (d *DescendingDeque) Push(num int) {
	for d.Len() > 0 && (*d)[d.Len()-1] < num {
		*d = (*d)[:d.Len()-1]
	}
	*d = append(*d, num)
}

func (d *DescendingDeque) Pop(num int) {
	if d.Len() > 0 && (*d)[0] == num {
		*d = (*d)[1:]
	}
}

func (d *DescendingDeque) Max() int {
	return (*d)[0]
}

func longestSubarray(nums []int, limit int) int {
	res := 0
	ascDeque, descDeque := &AscendingDeque{}, &DescendingDeque{}
	left := 0
	for right, num := range nums {
		ascDeque.Push(num)
		descDeque.Push(num)
		for descDeque.Max()-ascDeque.Min() > limit {
			ascDeque.Pop(nums[left])
			descDeque.Pop(nums[left])
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

/**
思路：
当滑动窗口移动时，需要用一个单调递增队列维护最小值，一个单调递减队列维护最大值。
若两个极值相减能满足约束条件，则表示其他元素两两相减必然更小，符合要求。

最开始只用一个单调递增队列维护最小值，用新元素与其相减看绝对值是否 <= limit。但未考虑到滑动窗口中的最大值与新元素相减的绝对值可能超过limit。
如limit = 2，滑动窗口 = [1, 2]，单调递增队列队口元素 = 1，新元素 = -1。
*/
