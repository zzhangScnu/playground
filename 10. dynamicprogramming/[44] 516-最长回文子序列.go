package dynamicprogramming

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] != s[j] {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				continue
			}
			if j-i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j-1] + 2
			}
		}
	}
	return dp[0][n-1]
}

/**
思路：
回文类题型，思路都是用二维数组承接s[i...j]范围内的计算结果，逐步递推。


DP数组及下标含义：
- i：s子字符串的起始位置；
- j：s子字符串的结束位置；
- dp[i][j]：s[i...j]中最长回文子串的长度。
由定义可知，i <= j。


递推公式：
情况一：s[i] == s[j]
s[i...j]是回文子串，那么dp[i][j] = dp[i+1][j-1] + 2
即在掐头去尾s[i+1...j-1]中的最长回文子串的长度的基础上，再加上s[i]和s[j]的长度。

这里需要额外考虑，当j-i==0，即i==j时，此时s[i...j]是一个字符，也是回文子串，长度为1。
这是一种边界，如果不加额外的分支，会导致dp[i][j] = dp[i+1][j-1] + 2 违反 i <= j的定义。

情况二：s[i] != s[j]
s[i...j]不是回文子串，那么dp[i][j]应该从两种情况中取最大值：
dp[i+1][j]：不考虑s[i]，即s[i+1...j]中最长回文子串的长度；
dp[i][j-1]：不考虑s[j]，即s[i...j-1]中最长回文子串的长度。

最终返回dp[0][n-1]，即s[0...n-1]中最长回文子串的长度。


初始化：
无需初始化，保持默认零值即可。


遍历方向：
因为dp[i][j]是由dp[i+1][j-1]推导出来的，即从左下方 -> 右上方。
所以遍历顺序应该是从下到上，从左到右。
*/


/**
两种实现方式对比：
第一种：无需额外初始化 & 需要额外分支
for i := n - 1; i >= 0; i-- {
	for j := i; j < n; j++ { // j从i开始，包含了单个字符的场景
		if s[i] != s[j] {
			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			continue
		}
		if j-i == 0 { // 这里需要对i == j即单个字符的场景开辟一条分支，以遵循i <= j的定义
			dp[i][j] = 1
		} else {
			dp[i][j] = dp[i+1][j-1] + 2
		}
	}
}

第二种：需要额外初始化 & 无需额外分支
for i := 0; i < n; i++ { // 单个字符的场景
	dp[i][i] = 1
}
for i := n - 1; i >= 0; i-- {
	for j := i + 1; j < n; j++ { // j从i+1开始，不包含单个字符的场景
		if s[i] != s[j] {
			dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			continue
		}
		dp[i][j] = dp[i+1][j-1] + 2
	}
}
*/
