package dynamicprogramming

// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
// 示例 1：
//
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
//
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
//
// 提示：
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
func minDistance72(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[l1][l2]
}

/**
1.DP数组及下标含义：
- i：word1当前下标为i；
- j：word2当前下标为j；
- dp[i][j]：word1[0...i]和word2[0...j]需要操作几步才能相等。
这里既可以操作word1，也可以操作word2。


2.递推公式：
比较两个字符时，dp[i][j]的取值有以下情况：

相等时，当前字符串无需额外操作，最小操作次数 = 前一个字符的操作次数：
dp[i - 1][j - 1]

不相等时：
替换：在word1[0...i-1]和word2[0...j-1]的最小编辑次数的基础上，增加此次替换的操作，即+1：
dp[i - 1][j - 1] + 1

删除：
即为跳过不考虑当前字符，当前字符的状态等于前一个字符的已求解状态：
删除word1中的字符：dp[i - 1][j] + 1
删除word2中的字符：dp[i][j - 1] + 1

添加：
在word1中添加字符，等同于在word2中删除字符：dp[i][j - 1] + 1
在word2中添加字符，等同于在word1中删除字符：dp[i - 1][j] + 1

所以递推公式即为这几种情况取最小值。


3. 初始化：
因为状态转移依赖上方和左方的值，所以需要对第一行和第一列进行初始化。
初始化的最小编辑次数，即为自身缩减至空字符串时的步数。


4. 遍历方向：
由左到右。
*/

/**
如何记录步骤
维护一个ops数组，用于记录步骤。
在收集最短编辑距离dp[i][j]时，将采取的操作(替换/删除/插入)记录到ops[i][j]中。
最终反向遍历ops数组，由ops[m][n]回溯到dp[0][0]即可得到完整的操作序列。

不能从ops[0][0]开始向后寻找，因为每一步都可能有多个分支路径，无法确定哪个有效路径最终能到达[m][n]。。
所以需要从终点开始向起点寻找，才能保证找到[起点 -> 终点]的完整正确操作序列。
*/
