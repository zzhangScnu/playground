package tree

// 给你一个产品数组 products 和一个字符串 searchWord ，products 数组中每个产品都是一个字符串。
//
// 请你设计一个推荐系统，在依次输入单词 searchWord 的每一个字母后，推荐 products 数组中前缀与 searchWord 相同的最多三个产品
// 。如果前缀相同的可推荐产品超过三个，请按字典序返回最小的三个。
//
// 请你以二维列表的形式，返回在输入 searchWord 每个字母后相应的推荐产品的列表。
//
// 示例 1：
//
// 输入：products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord
// = "mouse"
// 输出：[
// ["mobile","moneypot","monitor"],
// ["mobile","moneypot","monitor"],
// ["mouse","mousepad"],
// ["mouse","mousepad"],
// ["mouse","mousepad"]
// ]
// 解释：按字典序排序后的产品列表是 ["mobile","moneypot","monitor","mouse","mousepad"]
// 输入 m 和 mo，由于所有产品的前缀都相同，所以系统返回字典序最小的三个产品 ["mobile","moneypot","monitor"]
// 输入 mou， mous 和 mouse 后系统都返回 ["mouse","mousepad"]
//
// 示例 2：
//
// 输入：products = ["havana"], searchWord = "havana"
// 输出：[["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
//
// 示例 3：
//
// 输入：products = ["bags","baggage","banner","box","cloths"], searchWord =
// "bags"
// 输出：[["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"]
// ,["bags"]]
//
// 示例 4：
//
// 输入：products = ["havana"], searchWord = "tatiana"
// 输出：[[],[],[],[],[],[],[]]
//
// 提示：
//
// 1 <= products.length <= 1000
// 1 <= Σ products[i].length <= 2 * 10^4
// products[i] 中所有的字符都是小写英文字母。
// 1 <= searchWord.length <= 1000
// searchWord 中所有字符都是小写英文字母。
func suggestedProducts(products []string, searchWord string) [][]string {
	trie := TrieConstructor()
	for _, product := range products {
		trie.Insert(product)
	}
	var search func(prefix string, node *TrieNode) []string
	search = func(prefix string, node *TrieNode) []string {
		for _, ch := range prefix {
			if node.Children[ch-'a'] == nil {
				return []string{}
			}
			node = node.Children[ch-'a']
		}
		var res []string
		var traverse func(node *TrieNode, currentStr string)
		traverse = func(node *TrieNode, currentStr string) {
			if len(res) == 3 {
				return
			}
			if node.IsEndOfWord {
				res = append(res, currentStr)
			}
			for i, child := range node.Children {
				if child != nil {
					traverse(child, currentStr+string(rune('a'+i)))
				}
			}
		}
		traverse(node, prefix)
		return res
	}
	var res [][]string
	var prefix string
	for _, ch := range searchWord {
		prefix += string(ch)
		res = append(res, search(prefix, trie.root))
	}
	return res
}

/**
思路：
1. 用产品列表构建前缀树；
2. 将输入拆成不同长度的前缀，模拟逐个字符输入的情况；
3. 使用各前缀进行匹配。
   如果前缀匹配成功，意味着当前游标在前缀末尾的叶子节点上。此时需通过DFS遍历子树并收集结果。
   方法的参数中：
   node *TrieNode：当前搜索的前缀树节点，用于判断是否到达产品名称末尾；
   currentStr string：当前收集的产品名称路径，最终需加入结果集。
   因为前缀树的组织形式是alphabet的，因此输出结果是天然有序的。

注意，因为golang中'a'形式的rune可以直接隐式转换为uint8，因此node.Children[ch]可以通过编译；
但由前缀树的构建规则可知，node.Children[ch-'a']才是真正的[0, 25]区间范围的合法索引。
*/
