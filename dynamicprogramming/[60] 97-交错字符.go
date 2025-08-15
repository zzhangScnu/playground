package dynamicprogramming

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//
// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//
// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//
// 注意：a + b 意味着字符串 a 和 b 连接。
//
// 示例 1：
//
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
// 输出：true
//
// 示例 2：
//
// 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
// 输出：false
//
// 示例 3：
//
// 输入：s1 = "", s2 = "", s3 = ""
// 输出：true
//
// 提示：
//
// 0 <= s1.length, s2.length <= 100
// 0 <= s3.length <= 200
// s1、s2、和 s3 都由小写英文字母组成
//
// 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	l1, l2 := len(s1)+1, len(s2)+1
	dp := make([][]bool, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]bool, l2)
	}
	dp[0][0] = true
	for i := 1; i < l1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j < l2; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1] || dp[i][j-1] && s2[j-1] == s3[i+j-1]

		}
	}
	return dp[l1-1][l2-1]
}

/**
DP数组及下标含义
- i：s1当前指针；
- j：s2当前指针；
因为允许空字符串的存在，为了表达空，令i / j == 0时表示空值而不是第0位字符。所以dp数组大小需初始化为(len(s1) * len(s2))。
- dp[i][j]：s1[0, i - 1]和s2[0, j - 1]的组合是否能交错组成s3[0, i + j - 1]。

s1[0, i - 1]：i个字符
s2[0, j - 1]：j个字符
应组成s3中的i + j个字符，即s3[0, i + j - 1]。


递推公式
dp[i][j] =
	dp[i - 1][j] && s1[i - 1] == s3[i + j - 1] // s1[0, i - 1]和s2[0, j]能组成s3[0, i + j - 1]且本轮s1当前字符能与s3当前字符匹配
	||
	dp[i][j - 1] && s2[j - 1] == s3[i + j - 1] // s1[0, i]和s2[0, j - 1]能组成s3[0, i + j - 1]且本轮s2当前字符能与s3当前字符匹配
这里涵盖了仅s1与s3匹配、仅s2与s3匹配、s1和s2均与s3匹配任选一个的场景。


初始化
dp[0][0] = true
dp[0...len(s1)][0] -> 因s2为空字符串，逐个比对s1和s3中的字符匹配情况即可
dp[0][0...len(s2)] -> 因s1为空字符串，逐个比对s2和s3中的字符匹配情况即可

遍历方向
由递推公式可知，i / j 由 i - 1 / j - 1推导而来，
故由上而下，由左往右遍历。
*/
