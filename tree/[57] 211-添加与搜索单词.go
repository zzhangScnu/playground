package tree

// 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
//
// 实现词典类 WordDictionary ：
//
//
// WordDictionary() 初始化词典对象
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
// bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些'.' ，每个 . 都可以表示任何一个字母。
//
//
//
//
// 示例：
//
//
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search",
//"search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//
//解释：
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // 返回 False
//wordDictionary.search("bad"); // 返回 True
//wordDictionary.search(".ad"); // 返回 True
//wordDictionary.search("b.."); // 返回 True
//
//
// 提示：
//
//
// 1 <= word.length <= 25
// addWord 中的 word 由小写英文字母组成
// search 中的 word 由 '.' 或小写英文字母组成
// 最多调用 10⁴ 次 addWord 和 search

type WordDictionaryNode struct {
	Children    []*WordDictionaryNode
	IsEndOfWord bool
}

type WordDictionary struct {
	Root *WordDictionaryNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{
		Root: &WordDictionaryNode{
			Children: make([]*WordDictionaryNode, 26),
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.Root
	for _, c := range word {
		index := c - 'a'
		if cur.Children[index] == nil {
			cur.Children[index] = &WordDictionaryNode{
				Children: make([]*WordDictionaryNode, 26),
			}
		}
		cur = cur.Children[index]
	}
	cur.IsEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.doSearch(word, this.Root)
}

func (this *WordDictionary) doSearch(word string, cur *WordDictionaryNode) bool {
	if cur == nil {
		return false
	}
	if len(word) == 0 {
		return cur.IsEndOfWord
	}
	c := word[0]
	remainWord := word[1:]
	if c == '.' {
		for _, child := range cur.Children {
			if this.doSearch(remainWord, child) {
				return true
			}
		}
		return false
	}
	return this.doSearch(remainWord, cur.Children[c-'a'])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// todo: 如果第一个非空子节点不匹配，即使其他子节点可能匹配，也会返回 false 当所有子节点都为 nil 时，循环不会执行，导致遗漏返回 false 的逻辑
// todo: 迭代和递归混用
