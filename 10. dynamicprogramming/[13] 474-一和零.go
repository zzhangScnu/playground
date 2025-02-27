package dynamicprogramming

import "strings"

// 给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
//
// 请你找出并返回 strs 的最大子集的大小，该子集中 最多 有 m 个 0 和 n 个 1 。
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
//
// 示例 1：
//
// * 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
// * 输出：4
//
// * 解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
// 其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于 n 的值 3 。
//
// 示例 2：
// * 输入：strs = ["10", "0", "1"], m = 1, n = 1
// * 输出：2
// * 解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		cnt0, cnt1 := strings.Count(str, "0"), strings.Count(str, "1")
		for j := m; j >= cnt0; j-- {
			for k := n; k >= cnt1; k-- {
				dp[j][k] = max(dp[j][k], dp[j-cnt0][k-cnt1]+1)
			}
		}
	}
	return dp[m][n]
}

/**
二进制字符串数组，意味着只有'0'和'1'的组合。
要求【最多有m个0和n个1】的最大子集，表示子集中字符串总和的长度最大为0 -> m, 1 -> n。
故可以转化为0/1背包问题：
数组中的每一个字符串作为一个物品，选择取或不取，且最多只能选取一次。
背包容量有两个维度，分别表达0和1的容量。
当背包容量为m和n时，输出背包中能装下的最多的物品。

DP数组及下标含义：
- i：字符串候选集中对字符串i任意取/不取的情况，[0, len(strs)-1]；
- j：背包用于装0的容量为j，[0, m]；
- k：背包用于装1的容量为k，[0, n]；
- dp[i][j][k]：对字符串i取/不取且背包容量为j和k时，背包中满足总和最多有j个0和k个1的字符串的最大数量。

递推公式：
dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-len(strs[i]中的0)][k-len(strs[i]中的1)] + 1)

初始化：
dp[0][0][0] = 0：当背包容量能装下0个0和0和1时，字符串0无法放入，故最大数量为0。
其他单元格的值都可以推导出来。

遍历方向：从左到右，由上而下。
*/

/**
由上简化掉物品维度，只用一个二维数组来表达背包情况：
DP数组及下标含义：
- j：背包用于装0的容量为j，[0, m]；
- k：背包用于装1的容量为k，[0, n]；
- dp[j][k]：背包容量为j和k时，背包中满足总和最多有j个0和k个1的字符串的最大数量。

递推公式：
dp[j][k] = max(dp[j][k], dp[j-len(strs[i]中的0)][k-len(strs[i]中的1)] + 1)

初始化：
dp[0][0] = 0：当背包容量能装下0个0和0和1时，能放入字符串的最大数量为0。
其他单元格的值都可以推导出来。

遍历方向：
先物品，再背包，且背包倒序遍历；
背包的两个容量维度的遍历顺序可颠倒，不影响结果。
*/
