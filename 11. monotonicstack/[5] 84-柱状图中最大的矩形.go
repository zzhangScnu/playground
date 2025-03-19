package monotonicstack

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 示例 1:
//
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10
//
// 示例 2：
//
// 输入： heights = [2,4]
// 输出： 4
//
// 提示：
//
// 1 <= heights.length <=10⁵
// 0 <= heights[i] <= 10⁴
func largestRectangleArea(heights []int) int {
	heights = append(append([]int{0}, heights...), 0)
	st := []int{0}
	var res int
	for i := 1; i < len(heights); i++ {
		if heights[st[len(st)-1]] < heights[i] {
			st = append(st, i)
		} else if heights[st[len(st)-1]] == heights[i] {
			st = st[:len(st)-1]
			st = append(st, i)
		} else {
			for len(st) > 0 && heights[st[len(st)-1]] > heights[i] {
				mid, right := st[len(st)-1], i
				st = st[:len(st)-1]
				if len(st) > 0 {
					left := st[len(st)-1]
					area := (right - left - 1) * heights[mid]
					res = max(res, area)
				}
			}
			st = append(st, i)
		}
	}
	return res
}

/**
思路：
设基准柱子为j，高度为hj。j跟相邻的一系列柱子能组成的最大矩形：
向左寻找[0...j-1]区间内，高度小于j的柱子i；
向右寻找[j+1...len(height)-1]区间内，高度小于j的柱子k；

能跟j一起参与组成矩阵的柱子，一定是高度 >= hj的，否则某些柱子上方会留有空余，无法填满矩阵 -> 高度 = hj；
i / k为高度 < hj的柱子，无法参与以j为基准的矩阵面积计算，只取中间部分。即为左开右开区间 -> 宽度 = k - j - 1；
矩形面积 = (k - i - 1) * hj

为了寻找j左右两侧高度比j小的柱子，作为左开右开区间的端点，引入一个梯度递减的栈。

如heights = [1,5,6,2,3]，

栈底	<---梯度递减----	栈顶
st = [1, 5, 6
此时当前元素 = 2，栈口元素 = 6，
2和5正好包围了栈口元素6，即找到了一组i = 5，j = 6，k = 2。

【栈顶、栈顶的下一个元素、要入栈的三个元素 -> 要求最大面积的高度和宽度】

计算面积并尝试更新res后，弹出6，此时1和2正好包围了栈口元素5，即找到了一组i = 1，j = 5，k = 2。

计算面积并尝试更新res后，弹出5，此时栈内元素只剩1，不足下一次计算，将2入栈。

st = [1, 2, 3
此时元素遍历完了，但刚好是递减顺序，无法触发弹出进行计算。
【所以height数组需要人为在最后增加0，来规避这种情况。】

类似地，对于[5, 4, 3, 2, 1]这种case，
st = [5,
j = 5，k = 4，但栈内不足提供i以参与计算，此时仅弹出5，入栈4。
下一个j = 4，k = 3，也无法计算。
直到数组遍历完成，都无法算出任何一个矩阵面积。
【height数组需要人为在最前增加0，来规避这种情况。】

同时，前后增加0，还可以兼容heights = [1]，结果 = 1。

为什么接雨水不用？
因为接雨水最左和最右的柱子本来就不接雨水，如果是单调递减或单调递增的情况，
无法形成凹槽，接到的雨水为0，符合题意。
*/
