package unionfindset

// 给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!
// =b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
//
// 只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。
//
// 示例 1：
//
// 输入：["a==b","b!=a"]
// 输出：false
// 解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
//
// 示例 2：
//
// 输入：["b==a","a==b"]
// 输出：true
// 解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。
//
// 示例 3：
//
// 输入：["a==b","b==c","a==c"]
// 输出：true
//
// 示例 4：
//
// 输入：["a==b","b!=c","c==a"]
// 输出：false
//
// 示例 5：
//
// 输入：["c==c","b==d","x!=z"]
// 输出：true
//
// 提示：
//
// 1 <= equations.length <= 500
// equations[i].length == 4
// equations[i][0] 和 equations[i][3] 是小写字母
// equations[i][1] 要么是 '='，要么是 '!'
// equations[i][2] 是 '='
func equationsPossible(equations []string) bool {
	unionFindSet := NewUnionFindSet(26)
	for _, equation := range equations {
		left, right, operator := int(equation[0]-'a'), int(equation[3]-'a'), equation[1]
		if operator == '=' {
			unionFindSet.Union(left, right)
		}
	}
	for _, equation := range equations {
		left, right, operator := int(equation[0]-'a'), int(equation[3]-'a'), equation[1]
		if operator == '!' && unionFindSet.IsConnected(left, right) {
			return false
		}
	}
	return true
}

/**
思路：
1. 初始化并查集：
若运算符是 == ：将操作数连接到同一个联通分量中。
初始化后，相等操作符连接的元素将属于同一个联通分量。
2. 遍历 != 运算符连接的操作数：
若其属于同一个联通分量，则表示产生冲突，等式方程不可满足。

注意，需先遍历完 == 的方程完成初始化，再统一遍历 != 的方程。
不可同时遍历。
*/
