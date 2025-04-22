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

/**
组合：顺序无关，(1, 2) == (2, 1)
排列：顺序有关，(1, 2) != (2, 1)
*/

/**
递归问题三部曲：
1. 确定入参及返回值；
2. 确定base case即返回条件；
3. 确定单层处理逻辑。
*/

/**
回溯能解决什么问题？
回溯本质上也是暴力搜索，但能解决简单遍历不能解决的问题。
它适用于【考虑A的情况下的解】-【不考虑A(回溯)转而考虑B情况下的解】-【...】的问题，即在一棵涵盖所有可能解的树上遍历搜索答案的过程。
同时，回溯用递归来控制for循环嵌套的数量。
*/

/**
回溯问题可抽象为一棵N叉树：
树层：用for循环控制，本质是本次处理取候选集中的哪个元素；
树枝：用递归深度控制，本质是本次选取了某个元素后，后面结果的所有可能性；
回溯：在遍历树时不断调整结果。
*/

/**
回溯的剪枝操作：
将某些明显不可能存在答案的分支排除，可以减少遍历次数。
一般是在树层上控制，即for循环中的终止条件。
*/
