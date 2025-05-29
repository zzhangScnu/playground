package graph

import "container/list"

// 字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列
// beginWord -> s1 -> s2 -> ... -> sk：
//
// 每一对相邻的单词只差一个字母。
//
// 对于 1 <= i <= k 时，每个
// si 都在
// wordList 中。注意， beginWord 不需要在
// wordList 中。
//
// sk == endWord
//
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
// 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
//
// 示例 1：
//
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
// "log","cog"]
// 输出：5
// 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
//
// 示例 2：
//
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
// "log"]
// 输出：0
// 解释：endWord "cog" 不在字典中，所以无法进行转换。
//
// 提示：
//
// 1 <= beginWord.length <= 10
// endWord.length == beginWord.length
// 1 <= wordList.length <= 5000
// wordList[i].length == beginWord.length
// beginWord、endWord 和 wordList[i] 由小写英文字母组成
// beginWord != endWord
// wordList 中的所有字符串 互不相同
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{})
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	queue := list.New()
	queue.PushBack(beginWord)
	visited := make(map[string]int)
	visited[beginWord] = 1
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		from, to := element.Value.(string), []byte(element.Value.(string))
		if from == endWord {
			return visited[from]
		}
		for i := 0; i < len(from); i++ {
			origin := to[i]
			for ch := 'a'; ch <= 'z'; ch++ {
				to[i] = byte(ch)
				if _, ok := wordSet[string(to)]; ok && visited[string(to)] == 0 {
					queue.PushBack(string(to))
					visited[string(to)] = visited[from] + 1
				}
			}
			to[i] = origin
		}
	}
	return 0
}

/**
将所有的可能节点想象成一个无向有环图，需要在其上寻找一条起点到终点的最短路径。
使用BFS遍历方式，由其由内到外的扩散性可知，当寻找到终点时，路径必然最短。

不需要提前生成完整的图，而是边遍历边生成图的分支，扩大图的范围，直至找到终点。
记得遍历替换字符结束后，重置字符串。

需要引入辅助数组visited，作用有二：
1. 避免节点重复访问；
2. 记录起点到当前节点的最短路径长度。
*/

/**
注意：
Go中没有set内置实现，所以要通过map自行手搓一个。
其中value的数据类型是struct{}：
- 设置值时写入struct{}{}；
- 判断是否存在时通过if _, ok := wordSet[endWord]; ok。
*/
