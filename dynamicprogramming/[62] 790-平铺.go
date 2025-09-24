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
	var combineState func(t1, t2 bool) int
	combineState = func(t1, t2 bool) int {
		s1, s2 := 0, 0
		if t1 {
			s1 = 1
		}
		if t2 {
			s2 = 1
		}
		return s1<<1 | s2
	}
	var traverse func(i int, t1, t2 bool) int
	traverse = func(i int, t1, t2 bool) int {
		if i == n {
			return 1
		}
		if count, ok := dp[[2]int{i, combineState(t1, t2)}]; ok {
			return count
		}
		var count int
		t3, t4 := i+1 < n, i+1 < n
		if t1 && t2 {
			if t3 && t4 {
				count += traverse(i+1, false, false)
			}
			if t3 {
				count += traverse(i+1, false, true)
			}
			if t4 {
				count += traverse(i+1, true, false)
			}
		}
		if t1 && !t2 {
			if t3 && t4 {
				count += traverse(i+1, false, true)
			}
			if t3 {
				count += traverse(i+1, false, true)
			}
		}
		if !t1 && t2 {
			count += traverse(i+1, true, false)
		}
		count %= mod
		dp[[2]int{i, combineState(t1, t2)}] = count
		return count
	}
	return traverse(0, true, true)
}

/**
若要解决此题，需要结合当前列的摆放方式 + 下一列的可摆放情况，推导 I 形 和 L 形可能的摆放方式数。
定义一个矩阵，
|t1  t3|
|t2  t4|
其中，t1 和 t2 表示当前列的可用情况，若为 1 则表示可放置，无论是 I 形或 L 形。
t3 和 t4 表示下一列可选的摆放情况。
所以可以通过 t1 + t2 + t3 + t4 的状态，推导可能的摆放方式数。


DP数组及下标含义
- dp[i][t1][t2]：
- i：第 i 列
- t1：第 i 列中，纵向第 1 个方格是否可填充
- t2：第 i 列中，纵向第 2 个方格是否可填充


递推公式 // todo
经过 i - 1 列的摆放后，对于当前列 i 的 t1 和 t2，共有 3 种可能。
在这 3 种可能之上，结合 t3 和 t4是否可摆放（是否越界），共有 7 种不同的摆放情况：

i 的已摆放情况       	  i 的可用情况(t1 & t2)      	     i & i + 1 可能的摆放情况               i + 1 可用情况(t3 & t4)   					解析

0						   1							 1  1       1  0       1  1  		  0       1       0
0						   1	                         1  1       1  1       1  0    		  0       0       1							此时可以横向/垂直摆放两个 I 形，或以不同角度摆放一个 L 形

0                          1    						 1  1       0  1                 	  0       0	 								此时可以横向摆放一个 I 形，或摆放一个 L 形
1                          0							 1  0       1  1 					  1       0

1                          0 	  			 			 1  0       1  1                 	  1 	  0	   								此时可以横向摆放一个 I 形，或摆放一个 L 形
0                          1			     		     1  1       1  1                 	  0		  0									可以看出，这种情况跟上一种情况是对称的

1                          0				 			 1  1      1  0              		  0       1         						此时可以以不同方式摆放一个 L 形，
1                          0				 			 1  1      1  1                       0	 	  0		  							或以不同方式摆放两个 I 形，或仅在 i 列摆放一个 I 形。

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
