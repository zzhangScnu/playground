package dynamicprogramming

import "math"

// 给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
//
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
//
// 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎
// ，则可以在之后的操作中 重复使用 这枚鸡蛋。
//
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
//
// 示例 1：
//
// 输入：k = 1, n = 2
// 输出：2
// 解释：
// 鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
// 否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
// 如果它没碎，那么肯定能得出 f = 2 。
// 因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
//
// 示例 2：
//
// 输入：k = 2, n = 6
// 输出：3
//
// 示例 3：
//
// 输入：k = 3, n = 14
// 输出：4
//
// 提示：
//
// 1 <= k <= 100
// 1 <= n <= 10⁴
func superEggDropII(k int, n int) int {
	memo := make(map[[2]int]int)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == 0 || j == 0 {
			return 0
		}
		if i == 1 {
			return j
		}
		if j == 1 {
			return 1
		}
		if res, ok := memo[[2]int{i, j}]; ok {
			return res
		}
		res := math.MaxInt
		low, high := 1, j
		var mid int
		for low <= high {
			mid = i + (j-i)>>1
			broken, unbroken := dp(i-1, mid-1)+1, dp(i, j-mid)+1
			if broken > unbroken {
				res = min(res, unbroken)
				mid = high - 1
			} else {
				res = min(res, broken)
				mid = low + 1
			}
			memo[[2]int{low, high}] = res
		}
		return res
	}
	return dp(k, n)
}
