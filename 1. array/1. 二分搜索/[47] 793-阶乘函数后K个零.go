package ___二分搜索

// f(x) 是 x! 末尾是 0 的数量。回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 。
//
// 例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。
//
// 给定 k，找出返回能满足 f(x) = k 的非负整数 x 的数量。
//
// 示例 1：
//
// 输入：k = 0
// 输出：5
// 解释：0!, 1!, 2!, 3!, 和 4!均符合 k = 0 的条件。
//
// 示例 2：
//
// 输入：k = 5
// 输出：0
// 解释：没有匹配到这样的 x!，符合 k = 5 的条件。
//
// 示例 3:
//
// 输入: k = 3
// 输出: 5
//
// 提示:
//
// 0 <= k <= 10⁹
func preimageSizeFZF(k int) int {
	n := 5 * (k + 1)
	leftBound, rightBound := searchLeftBoundFZF(n, k), searchRightBoundFZF(n, k)
	return rightBound - leftBound + 1
}

func trailingZeroesFZF(num int) int {
	var res int
	for i := num; i > 0; i /= 5 {
		res += i / 5
	}
	return res
}

func searchLeftBoundFZF(n int, target int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)>>1
		if trailingZeroesFZF(mid) >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchRightBoundFZF(n int, target int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)>>1
		if trailingZeroesFZF(mid) <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}

/**
要求有多少个x`!，满足有k个尾随零。
x`! ↑ -> 尾随零数量↑，存在单调递增关系，故可尝试引入二分搜索求解。

三部曲：

1. 画出函数在二维坐标上的图像，明确 x、f(x)、target，并实现函数 f；
   - x：阶乘x`!中的x`。因为x`和x`-1...分别能表示不同的阶乘，且本题求的是有多少个这样的阶乘，故求解的是x的区间；
   - f(x)：尾随零的数量。
   - target：k。

2. 明确 x 的取值范围，作为二分搜索的搜索区间，初始化left和right变量；
	- left：0；
	- right：由于x`!中的k个尾随零，由阶乘中的k个因数5提供，所以x的区间应为[5 * k, 5 * (k + 1))。

3. 根据题意明确使用搜索左侧 / 右侧的二分搜索算法，写出解法代码。
	f(x) == k的约束下，求x的左边界和右边界，相减即为所求数量。


值得注意的是边界条件的处理被隐含了：
如果没有满足k的x`!，如k = 5，此时右边界right = 30，
leftBound = 25, rightBound = 24,
res = 0，满足题意。
*/
