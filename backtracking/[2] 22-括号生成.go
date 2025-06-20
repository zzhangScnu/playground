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

/**
思路：回溯法
在生成括号&回溯过程中，分别维护左括号和右括号的剩余额度。
- base case1：
  因为是以从左到右的顺序生成括号，所以合法情况下任一时刻的左括号数量 >= 右括号数量，
  即左括号剩余额度 <= 右括号剩余额度；
- base case2：
  超出题目给的n的数量限制；
- base case3：
  若当前左右括号的剩余额度都到达了0，且没有命中上述不合法的base case，则表示生成了一个合法括号串。
  此时收集结果到res中。
  收集完成后，无需手动将path置空，因为算法本身带有回溯逻辑。

回溯时，分别对左括号和右括号进行生成和回溯。

第一版的思路是，没有base case做在途检查，而是生成了长度为2n的括号串后，
再用栈辅助的方式进行统一检查，过滤掉非法值。这样效率较低。
*/
