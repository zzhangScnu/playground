package tree

// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使
// 用。
//
// 示例 1：
//
// 输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f",
// "l","v"]], words = ["oath","pea","eat","rain"]
// 输出：["eat","oath"]
//
// 示例 2：
//
// 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
// 输出：[]
//
// 提示：
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 10⁴
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同
func findWords(board [][]byte, words []string) []string {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	trie := TrieConstructor()
	for _, word := range words {
		trie.Insert(word)
	}
	var res []string
	var traverse func(x, y int, curNode *TrieNode, path []byte)
	traverse = func(x, y int, curNode *TrieNode, path []byte) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		ch := board[x][y]
		if ch == ' ' {
			return
		}
		if curNode.Children[ch-'a'] == nil {
			return
		}
		path = append(path, ch)
		board[x][y] = ' '
		defer func() {
			board[x][y] = ch
			path = path[:len(path)-1]
		}()
		curNode = curNode.Children[ch-'a']
		if curNode.IsEndOfWord {
			res = append(res, string(path))
			curNode.IsEndOfWord = false
		}
		for _, movement := range movements {
			traverse(x+movement[0], y+movement[1], curNode, path)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			traverse(i, j, trie.root, []byte{})
		}
	}
	return res
}

//1. 核心问题：结果重复添加
//当 Trie 中存在「前缀包含关系」的单词（如 "a" 和 "aa"）时，遍历到长单词时会重复添加短单词。
//例如：找到 "aa" 时，由于 "a" 的 IsEndOfWord 为 true，会在遍历第一个 'a' 时添加 "a"，遍历第二个 'a' 时再次添加 "a"，导致最终结果包含重复字符串（如 ["a", "a", "aa"]）。
//解法：set收集结果 或 魔改前缀树节点

/**
思路：
1. 共享前缀路径，剪枝冗余扩展方向；
2. 递归函数维护几个当前变量：元素坐标、前缀树节点、路径结果集；
3. 正常使用回溯维护访问标记和路径结果集；
4. 当找到一条可行路径时，在收集结果时，将前缀树的节点的 IsEndOfWord 置为 false，避免出现重复结果；
   原因：当 Trie 中存在「前缀包含关系」的单词（如 "a" 和 "aa"）时，遍历到长单词时会重复添加短单词。
   例如：找到 "aa" 时，由于 "a" 的 IsEndOfWord 为 true，会在遍历第一个 'a' 时添加 "a"，遍历第二个 'a' 时再次添加 "a"，导致最终结果包含重复字符串（如 ["a", "a", "aa"]）。
   解法：使用集合收集结果 或 魔改前缀树节点。
5. 对于每个元素进行路径探索。
*/
