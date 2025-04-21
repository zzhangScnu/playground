package multiplepointers

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
//
// 示例 1：
//
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49
// 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
//
// 示例 2：
//
// 输入：height = [1,1]
// 输出：1
//
// 提示：
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
func maxArea(height []int) int {
	var res int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			res = max(res, height[left]*(right-left))
			left++
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
	}
	return res
}

/**
思路：

与接雨水类似，但不同点在于，接雨水需要针对每根柱子能够承接的雨水分别进行计算再加总，
而本题仅需针对某两根柱子围出的面积求取最大值，即采用二分搜索的思想，
根据左右柱子的相对高度关系，分别动态计算面积取优，并推进选取下一根柱子，与另一根柱子共同围成新的矩形。
直至两根柱子重叠，即无法围成一个矩形时结束。

为什么不需要两两柱子匹配从中取优的O(n^2)暴力解法，而是可以优化为O(n)线性解法？
1. 采取剪枝策略：
对于柱子left和right来说，围成的矩阵面积为min(height[left], height[right]) * (right - left)。
其中取min的原因是短板效应，即矩阵面积只能由最矮的柱子决定。
- 若此时height[left] < height[right]，说明此时矩阵面积为height[left] * (right - left)。
	- 此时不考虑left和[left + 1, right)区间内的柱子围成的矩形的原因是：
		- 高度：依然取决于min(height[left], height[left + 1, right])，即不可能比当前的height[left]更高；
		- 宽度：right - left，[left + 1, right)区间内的柱子的索引，比right更小。
		- 综上，left和[left + 1, right)区间内的柱子围成的矩形面积必然更小，故可以直接剪枝；
	- 此时推进left，即left++，目的是在贪心策略下，寻求更高的柱子left，试图令min(height[left], height[right]) * (right - left)更大。
- 同理，对于height[right] < height[left]的情况也一样。
2. 为什么可以由两边向中间夹逼？
因为我们只需要不断求取矩形的最大面积，而这种做法不会遗漏更大面积的场景。
在固定left和right且取得截至目前的最大矩形面积后：
- 若固定left，right--，则由短板效应可知高度不会超过height[left]，而宽度right-left会减小，所以一定不会出现更大的面积；
- 若固定right，left++，虽然宽度right-left会减小，但可能会找到更大的height[left]，可以取min(height[left], height[right])作为更大的高度，有可能找到更大的面积。
同时，双指针每次只移动一个指针而不是两个，确保每次移动都是基于当前的最优选择，不会跳过可能的更大组合。
每次移动后剩余的可能解空间仍然包含最优解。假设当前的最优解存在于未检查的组合中，那么移动较矮的一边后，最优解仍然存在于剩余的组合中，因为被排除的组合的水量都不可能超过当前的最大值。
*/
