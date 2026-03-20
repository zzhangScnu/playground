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

/**
遍历字符可能性的实现：

实现一：
for ch := 'a'; ch <= 'z'; ch++ {
	// ...
}
ch 为 int32

实现二：
for i := range 26 { // Go 1.22 及 以后支持这种语法
	newChar := rune('a' + i)
	// ...
}
*/

/*
*
- 用 wordMap 代替 visited
- 用 [][]interface{}{{word, step}} 代替 container.list（BFS 队列） + visited（重复性校验 + 步数记录）
*/
func ladderLengthII(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]interface{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	queue := [][]interface{}{{beginWord, 1}}
	for len(queue) > 0 {
		cur, step := queue[0][0].(string), queue[0][1].(int)
		queue = queue[1:]
		curChars := []rune(cur)
		for pos, char := range curChars {
			tmpChar := char
			for i := 0; i < 26; i++ {
				newChar := rune('a' + i)
				if newChar == tmpChar {
					continue
				}
				curChars[pos] = newChar
				next := string(curChars)
				if next == endWord {
					return step + 1
				}
				if _, ok := wordMap[next]; !ok {
					continue
				}
				delete(wordMap, next)
				queue = append(queue, []interface{}{next, step + 1})
			}
			curChars[pos] = tmpChar
		}
	}
	return 0
}

/**
最开始的实现，想着节省内存，将变量的定义放在了循环外，但是发现这样优化不了多少：

1. 基本类型（string/int/rune）：几乎无节省
cur/next（string）：Go 中 string 是不可变的，cur = queue[0][0].(string) 只是让cur指向新的字符串地址，没有内存分配，复用和重新声明的开销完全一样；
step（int）：赋值只是修改一个 8 字节的内存值，循环内声明也只是在栈上分配 8 字节，栈内存分配的开销是纳秒级（几乎感知不到）；
tmpRune（rune）：4 字节的栈内存，同理无差异。

2. 引用类型（[] rune）：省的是 “切片头”，不是 “底层数组”
curChars = []rune(cur)：无论curChars是全局还是循环内声明，底层字符数组都会重新分配（因为每个单词的字符都不一样）；
唯一能省的是 “切片头”（一个包含指针、长度、容量的结构体，共 24 字节）—— 但 24 字节的栈内存分配，对现代 CPU 来说完全是 “毛毛雨”。

二、“省时间 / 空间” 的误区：栈内存 vs 堆内存
 “复用变量省空间”，核心是混淆了「栈内存」和「堆内存」：
栈内存（循环内声明的变量）：由编译器自动分配 / 释放，速度极快（纳秒级），且栈大小通常只有几 MB，24 字节的切片头完全不影响；
堆内存（比如wordMap、queue）：才是程序内存开销的大头，代码中用delete(wordMap, next)优化堆内存，这才是真正有意义的优化 —— 而纠结栈上的 24 字节，属于 “捡芝麻丢西瓜”。

三、什么时候 “复用变量” 才有意义？
只有满足以下条件，复用变量才值得做：
变量是堆内存分配的大对象（比如长度 10 万的切片、大 map）；
循环次数极多（比如 1 亿次以上）；
性能是核心指标（比如高频交易系统、实时渲染）。

而「单词接龙」的场景：
单词长度通常≤10，循环次数≤1000；
堆内存开销主要在queue和wordMap，栈内存几乎可以忽略；
开发效率（调试、维护）远比重微的性能更重要。
*/
