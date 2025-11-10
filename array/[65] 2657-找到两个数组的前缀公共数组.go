package array

// 给你两个下标从 0 开始长度为 n 的整数排列 A 和 B 。
//
// A 和 B 的 前缀公共数组 定义为数组 C ，其中 C[i] 是数组 A 和 B 到下标为 i 之前公共元素的数目。
//
// 请你返回 A 和 B 的 前缀公共数组 。
//
// 如果一个长度为 n 的数组包含 1 到 n 的元素恰好一次，我们称这个数组是一个长度为 n 的 排列 。
//
// 示例 1：
//
// 输入：A = [1,3,2,4], B = [3,1,2,4]
// 输出：[0,2,3,4]
// 解释：i = 0：没有公共元素，所以 C[0] = 0 。
// i = 1：1 和 3 是两个数组的前缀公共元素，所以 C[1] = 2 。
// i = 2：1，2 和 3 是两个数组的前缀公共元素，所以 C[2] = 3 。
// i = 3：1，2，3 和 4 是两个数组的前缀公共元素，所以 C[3] = 4 。
//
// 示例 2：
//
// 输入：A = [2,3,1], B = [3,1,2]
// 输出：[0,1,3]
// 解释：i = 0：没有公共元素，所以 C[0] = 0 。
// i = 1：只有 3 是公共元素，所以 C[1] = 1 。
// i = 2：1，2 和 3 是两个数组的前缀公共元素，所以 C[2] = 3 。
//
// 提示：
//
// 1 <= A.length == B.length == n <= 50
// 1 <= A[i], B[i] <= n
// 题目保证 A 和 B 两个数组都是 n 个元素的排列。
func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	prefix, res := make([]int, n+1), make([]int, n)
	var commonPrefixCount int
	for i := 0; i < n; i++ {
		a, b := A[i], B[i]
		prefix[a]++
		if prefix[a] == 2 {
			commonPrefixCount++
		}
		prefix[b]++
		if prefix[b] == 2 {
			commonPrefixCount++
		}
		res[i] = commonPrefixCount
	}
	return res
}

/**
思路：
维护一个表示"到目前为止，A 和 B 中公共字符的长度"的整型变量 commonPrefixCount。
维护一个表示"到目前为止，A 和 B 中各字符出现次数"的数组变量 prefix。
如果某一步中，令 prefix[字符] == 2，则将 commonPrefixCount++，并赋值 res[i] = commonPrefixCount，
表示迄今为止，A 和 B 中公共字符的长度增加。
最终 res 即为结果。
*/
