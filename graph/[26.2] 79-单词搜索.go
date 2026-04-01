package graph

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 示例 1：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCCED"
// 输出：true
//
// 示例 2：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "SEE"
// 输出：true
//
// 示例 3：
//
// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word =
// "ABCB"
// 输出：false
//
// 提示：
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
// 进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
func existII(board [][]byte, word string) bool {
	movements := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	trie := TrieConstructor()
	trie.Insert(word)
	var traverse func(x, y int, curNode *TrieNode) bool
	traverse = func(x, y int, curNode *TrieNode) bool {
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		ch := board[x][y]
		if ch == ' ' {
			return false
		}
		idx := getCharIndex(ch)
		if curNode.Children[idx] == nil {
			return false
		}
		board[x][y] = ' '
		curNode = curNode.Children[idx]

		if curNode.IsEndOfWord {
			return true
		}
		for _, movement := range movements {
			if traverse(x+movement[0], y+movement[1], curNode) {
				return true
			}
		}
		board[x][y] = ch
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if traverse(i, j, trie.root) {
				return true
			}
		}
	}
	return false
}

// A-Z(0-25), a-z(26-51)
func getCharIndex(c byte) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A')
	}
	return int(c - 'a' + 26)
}

type TrieNode struct {
	Children    []*TrieNode
	IsEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func TrieConstructor() Trie {
	return Trie{
		root: &TrieNode{
			Children: make([]*TrieNode, 52),
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		idx := getCharIndex(c)
		if cur.Children[idx] == nil {
			cur.Children[idx] = &TrieNode{
				Children: make([]*TrieNode, 52),
			}
		}
		cur = cur.Children[idx]
	}
	cur.IsEndOfWord = true
}

/**
用前缀树思路解决。

但是需要注意的一点是，字母集合里面是【a-zA-Z】，即范围为 52。
func getCharIndex(c byte) int {
	if c >= 'A' && c <= 'Z' { // 大写字母，占用[0, 25]
		return int(c - 'A')
	}
	return int(c - 'a' + 26) // 小写字母，占用[26, 51]
}
*/
