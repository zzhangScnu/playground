package array

// 珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
//
// 珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后
// 这一小时内不会再吃更多的香蕉。
//
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
//
// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。
//
// 示例 1：
//
// 输入：piles = [3,6,7,11], h = 8
// 输出：4
//
// 示例 2：
//
// 输入：piles = [30,11,23,4,20], h = 5
// 输出：30
//
// 示例 3：
//
// 输入：piles = [30,11,23,4,20], h = 6
// 输出：23
//
// 提示：
//
// 1 <= piles.length <= 10⁴
// piles.length <= h <= 10⁹
// 1 <= piles[i] <= 10⁹
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1_000_000_000
	for left <= right {
		mid := left + (right-left)/2
		t := time(piles, mid)
		if t == h {
			right = mid - 1
		} else if t < h {
			right = mid - 1
		} else if t > h {
			left = mid + 1
		}
	}
	return left
}

func time(piles []int, k int) int {
	var t int
	for _, num := range piles {
		t += num / k
		if num%k > 0 {
			t++
		}
	}
	return t
}

/**
思路：构建一个关于x的f(x)单调函数。
想象一个平面直角坐标系中，递增或递减的一个函数。
将题目转化为：当f(x)满足给定的target时，求x的极值。

三部曲：
1. 画出函数在二维坐标上的图像，明确 x、f(x)、target，并实现函数 f：
	- x = k，所求的吃香蕉速度；
	- f(x) =  吃掉所有香蕉的时间。隐含按顺序吃每堆香蕉。
	  速度↑ -> 时间↓，存在单调关系。若吃香蕉的速度 = x / 小时，则需要f(x)小时吃完。
	- target = h
	  吃香蕉的时间是对f(x)返回值的最大约束，需满足f(x) <= target。

2. 明确 x 的取值范围，作为二分搜索的搜索区间，初始化left和right变量：
    - left：1；
	- right：每小时最多吃一堆香蕉，故最大值为piles数组中的最大值；
		- 方式一：遍历piles计算最大值；
		- 方式二：取题目约束1 <= piles[i] <= 10^9。

3. 根据题意明确使用搜索左侧 / 右侧的二分搜索算法，写出解法代码。
	使用二分搜索，求x的左边界。
	函数是单调递减的，围绕中间点mid去收缩左右边界 -> 让f(x)在满足约束的情况下，x尽可能小。

当想不通的时候，把函数图像画出来吧=)
*/
