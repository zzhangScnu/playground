package backtracking

import (
	"strings"
)

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
//
// 提示：
//
// 1 <= n <= 8
func generateParenthesis(n int) []string {
	var path, res []string
	var doGenerateParenthesis func(leftRemain, rightRemain int)
	doGenerateParenthesis = func(leftRemain, rightRemain int) {
		if leftRemain > rightRemain {
			return
		}
		if leftRemain < 0 || rightRemain < 0 {
			return
		}
		if leftRemain == 0 && rightRemain == 0 {
			res = append(res, strings.Join(path, ""))
			return
		}
		path = append(path, "(")
		doGenerateParenthesis(leftRemain-1, rightRemain)
		path = path[:len(path)-1]
		path = append(path, ")")
		doGenerateParenthesis(leftRemain, rightRemain-1)
		path = path[:len(path)-1]
	}
	doGenerateParenthesis(n, n)
	return res
}
