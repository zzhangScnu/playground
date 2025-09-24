package dynamicprogramming

// 有两种形状的瓷砖：一种是 2 x 1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。
//
// 给定整数 n ，返回可以平铺 2 x n 的面板的方法的数量。返回对 10⁹ + 7 取模 的值。
//
// 平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形。
//
// 示例 1:
//
// 输入: n = 3
// 输出: 5
// 解释: 五种不同的方法如上所示。
//
// 示例 2:
//
// 输入: n = 1
// 输出: 1
//
// 提示：
//
// 1 <= n <= 1000
func numTilings(n int) int {
	mod := 1_000_000_007
	dp := make(map[[2]int]int)
	var combineState func(t1, t2 int) int
	combineState = func(t1, t2 int) int {
		return t1 << 1 | t2
	}
	var traverse func(i int, t1, t2 int) int
	traverse = func(i int, t1, t2 int) int {
		if i == n {
			return 1
		}
		if count, ok := dp[[2]int{i, combineState(t1, t2)}]; ok {
			return count
		}
		var count int
		if t1 == 0 && t2 == 0 {
			count += traverse(i+1, 0, 0)
		}
		if t1 == 0 && t2 == 1 {
			count += traverse(i+1, 0, 1) + traverse(i+1, 1, 1)
		}
		if t1 == 1 && t2 == 0 {
			count += traverse(i+1, 1, 0) + traverse(i+1, 1, 1)
		}
		if t1 == 1 && t2 == 1 {
			count += traverse(i+1, 0, 0) + traverse(i+1, 1, 0) + traverse(i+1, 0, 1) + traverse(i+1, 1, 1)
		}
		count %= mod
		dp[[2]int{i, combineState(t1, t2)}] = count
		return count
	}
	return traverse(0, 0, 0)
}

/**
若要解决此题，需要结合当前列的摆放方式，推演下一列的摆放方式数。
定义一个矩阵，
|t1  t3|
|t2  t4|
其中，t1 和 t2 表示当前列的占用情况，若为 1 则表示已摆放，无论是 I 形或 L 形。
t3 和 t4 表示下一列可选的摆放情况。
所以可以通过 t1 + t2 的状态，推导 t3 + t4 的状态。


DP数组及下标含义
- dp[i][t1][t2]：
- i：第 i 列
- t1：第 i 列中，纵向第 1 个方格填充的情况
- t2：第 i 列中，纵向第 2 个方格填充的情况


递推公式
对于当前列 i，t1 和 t2 有 4 种情况：

i 的情况                   i & i + 1 的情况                  				解析
0                         0  0                            				i - 1 可能竖向放置了一个 I 形。若 t1 和 t2 不摆放，即 t1 和 t2 均为0，
0                         0  0							  				此时因为不能跳过某些位置留空而进行下一个位置摆放的约束，t3 和 t4 只能为0


0                         0  0       0  1                 				此时可以横向摆放一个 I 形，或摆放一个 L 形
1                         1  1       1  1

1                         1  1       1  1                 				此时可以横向摆放一个 I 形，或摆放一个 L 形
0                         0  0       0  1                 				可以看出。这种情况跟上一种情况是对称的

1                         1  0       1  1      1  0      1  1    		此时可以以不同方式摆放一个 L 形，
1                         1  0       1  0      1  1      1  1           或以不同方式摆放两个 I 形，或仅在 i 列摆放一个 I 形。

综上，
dp[i][t1][t2] ->
	dp[i + 1][t1][t2]，i + 1 处的 t1 和 t2 即为 t3 和 t4。
因为三维数组的边界和初始化比较难处理，所以这题用自顶向下的递归方式解决更简单。
当 i == n 时，表示找到了一种合法的摆放方式。


初始化
因为用递归方式解决，所以做好base case编写即可。
此外，因为存在大量重叠子问题，所以需要引入备忘录模式。


遍历方向
从左到右，从上往下。
*/
