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
	var path []byte
	var traverse func(x, y, wordIndex, charIndex int)
	traverse = func(x, y, wordIndex, charIndex int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if charIndex == len(words[wordIndex]) {
			res = append(res, words[wordIndex])
			return
		}
		if board[x][y] == ' ' {
			return
		}
		path = append(path, board[x][y])
		board[x][y] = ' '
		defer func() {
			board[x][y] = path[len(path)-1]
			path = path[:len(path)-1]
		}()
		if !trie.StartsWith(string(path)) {
			return
		}
		for _, movement := range movements {
			traverse(x+movement[0], y+movement[1], wordIndex, charIndex+1)
		}
	}
	for k := 0; k < len(words); k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				traverse(i, j, k, 0)
			}
		}
	}
	return res
}
