package backtracking

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
//
// 示例 1：
//
// 输入：n = 4, k = 2
// 输出：
// [
//
//	[2,4],
//	[3,4],
//	[2,3],
//	[1,2],
//	[1,3],
//	[1,4],
//
// ]
//
// 示例 2：
//
// 输入：n = 1, k = 1
// 输出：[[1]]
//
// 提示：
//
// 1 <= n <= 20
// 1 <= k <= n

var combination []int

var combinations [][]int

func combine(n int, k int) [][]int {
	combination, combinations = []int{}, [][]int{}
	doCombine(1, n, k)
	return combinations
}

func doCombine(beginNum int, n int, k int) {
	if len(combination) == k {
		res := make([]int, k)
		copy(res, combination)
		combinations = append(combinations, res)
		return
	}
	for i := beginNum; i <= n-(k-len(combination))+1; i++ {
		combination = append(combination, i)
		doCombine(i+1, n, k)
		combination = combination[:len(combination)-1]
	}
}
