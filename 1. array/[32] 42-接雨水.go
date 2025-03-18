package array

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
//
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 提示：
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
func trap(height []int) int {
	var res int
	for i := 0; i < len(height); i++ {
		left, right, mid := height[i], height[i], height[i]
		for l := i - 1; l >= 0; l-- {
			if height[l] > left {
				left = height[l]
			}
		}
		for r := i + 1; r < len(height); r++ {
			if height[r] > right {
				right = height[r]
			}
		}
		res += min(left, right) - mid
	}
	return res
}

func trapII(height []int) int {
	size := len(height)
	left, right := make([]int, size), make([]int, size)
	left[0] = height[0]
	for i := 1; i < size; i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	var res int
	for i := 0; i < size; i++ {
		sum := min(left[i], right[i]) - height[i]
		res += sum
	}
	return res
}

func trapIII(height []int) int {
	var res int
	l, r, lMaxHeight, rMaxHeight := 0, len(height)-1, height[0], height[len(height)-1]
	for l <= r {
		lMaxHeight = max(lMaxHeight, height[l])
		rMaxHeight = max(rMaxHeight, height[r])
		if lMaxHeight < rMaxHeight {
			res += lMaxHeight - height[l]
			l++
		} else {
			res += rMaxHeight - height[r]
			r--
		}
	}
	return res
}

/**
思路：
可以将承接雨水的凹槽切割成宽度为1的格子，
则问题分解为：求每个柱子竖直方向上能接雨水的格子面积的总和，
进一步分解为：求每个柱子上方能接雨水的格子的高度。
(对应地，单调栈地解法，是求三个柱子围成的凹槽之间的横向方向上能接雨水的面积)


- 暴力解法
1. for循环遍历每一个柱子；
2. 对每一个柱子向两边遍历，寻找左侧最高和右侧最高的柱子；
3. 每个柱子上方能接雨水的格子的高度 = min(左侧最高的柱子高度, 右侧最高的柱子高度) - 当前的柱子高度。
柱子遍历范围为[0, len(height)-1]，且 左侧最高的柱子高度 和 右侧最高的柱子高度 初始化为当前柱子的高度，
所以能兼容最左侧柱子和最右侧柱子不接雨水的场景，代入公式相减为0，不影响结果。
时间复杂度为O(n2)，空间复杂度为O(1)。


- 备忘录优化
在暴力解法的基础上，通过预处理，先行找到每根柱子的 左侧最高的柱子高度 和 右侧最高的柱子高度，
在计算时直接索引取值，进行计算。
时间复杂度为O(n)，空间复杂度为O(n)。


- 双指针优化
个人感觉类似滚动数组的思想，不是将值预存下来，而是边计算边滚动更新。
跟备忘录有一点不同的地方在于，
备忘录是以当前柱子为基准，找其 左侧最高的柱子高度 和 右侧最高的柱子高度；
而双指针是从两边开始夹逼，找全局 左侧最高的柱子高度 和 右侧最高的柱子高度。
再对于较矮的一侧，进行计算。类似于木桶短板效应，能接多少水取决于较短处。
*/
