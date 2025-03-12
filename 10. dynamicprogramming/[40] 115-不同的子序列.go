package dynamicprogramming

// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10⁹ + 7 取模。
//
// 示例 1：
//
// 输入：s = "rabbbit", t = "rabbit"
// 输出：3
// 解释：
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// rabbbit
// rabbbit
// rabbbit
//
// 示例 2：
//
// 输入：s = "babgbag", t = "bag"
// 输出：5
// 解释：
// 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
// babgbag
// babgbag
// babgbag
// babgbag
// babgbag
//
// 提示：
//
// 1 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成
func numDistinct(s string, t string) int {
	sl, tl := len(s), len(t)
	dp := make([][]int, tl+1)
	for i := 0; i <= tl; i++ {
		dp[i] = make([]int, sl+1)
	}
	for j := 0; j <= sl; j++ {
		dp[0][j] = 1
	}
	for i := 1; i <= tl; i++ {
		for j := 1; j <= sl; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[tl][sl]
}

/**
思路：
跟392类似，在比较过程中，遇到不匹配的情况会有"删除"操作，即移动指针跳过当前字符。
注意本题已经隐含了t是s的子序列的前提条件。


举例：
t = bag
s = bagg


DP数组及下标含义：
- i：t(target)中的当前下标 == i - 1；
- j：s(source)中的当前下标 == j - 1；
- dp[i][j]：t[0...i - 1]在s[0...j - 1]中出现的个数，即s[0...j - 1]中有几种方式可以组成t[0...i - 1]。

递推公式：
if t[i - 1] == s[j - 1] // t中当前字符 == s中当前字符
	dp[i][j] = dp[i - 1][j - 1] // 组成个数继承之前字符串的，比如t由"ba" -> "ba[g]"，s由"bag" -> "bag[g]"
			 + dp[i][j - 1] // 如果t中当前字符 == s中前一个字符，比如t由"ba" -> "ba[g]"，s由"ba" -> "ba[g]"，此时依然能匹配，则说明是另一种组成方式

初始化：
因为主要的转移是dp[i][j] = dp[i - 1][j - 1]，所以保留第一行和第一列为零值是不够的。
需要额外处理。
dp[0][1...len(s)-1]：t == ""，s不为空，则有1种方式可以由s组成t；
dp[1...len(t)-1][0]：t不为空，s == ""，则有0种方式可以由s组成t；
dp[0][0]：t == ""，s == ""，则有1种方式可以由s组成t。

遍历方向：从左到右。
*/

/**
if t[i - 1] == s[j - 1] // t中当前字符 == s中当前字符
	dp[i][j] = dp[i - 1][j - 1] // 则个数继承之前字符串的，比如t由"ba" -> "ba[g]"，s由"bag" -> "bag[g]"
	if t[i - 1] == s[j - 2] // 如果t中当前字符 == s中前一个字符，比如t由"ba" -> "ba[g]"，s由"ba" -> "ba[g]"，此时依然能匹配，则说明是另一种组成方式
		dp[i][j] = dp[i - 1][j - 2] + 1 // 则组成个数要+1
else ...

之前递推公式写成这样，其实动态规划跟递归相似，
在本层只考虑本层的逻辑，其他的交由前序推导/递归处理即可。
*/
