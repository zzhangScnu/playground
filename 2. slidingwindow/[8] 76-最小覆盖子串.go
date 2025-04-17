package slidingwindow

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 示例 2：
//
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
//
// 示例 3:
//
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
// 进阶：你能设计一个在
// o(m+n) 时间内解决此问题的算法吗？
func minWindow(s string, t string) string {
	var i, validNum int
	targetCnt, windowCnt := make(map[string]int), make(map[string]int)
	for _, char := range t {
		targetCnt[string(char)]++
	}
	ri, rl := 0, 100_001
	for j := 0; j < len(s); j++ {
		rc := string(s[j])
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				validNum++
			}
			for validNum == len(targetCnt) {
				if j-i+1 < rl {
					ri = i
					rl = j - i + 1
				}
				lc := string(s[i])
				if targetCnt[lc] > 0 {
					if windowCnt[lc] == targetCnt[lc] {
						validNum--
					}
					windowCnt[lc]--
				}
				i++
			}
		}
	}
	if rl == 100_001 {
		return ""
	}
	return s[ri : ri+rl]
}

/**
思路：
- 滑动窗口：在s中截取且不断从左向右滑动的子字符串；
- i & j：滑动窗口的左 / 右边界索引；
- targetCnt：字符串t中的每个字符 -> 出现次数；
- windowCnt：滑动窗口中的每个字符 -> 出现次数；
- validNum：滑动窗口中符合t中出现要求的字符的个数。
  当validNum == len(targetCnt)时，表示t中所有字符均在滑动窗口中涵盖，且出现次数与t中一致，
  此时的滑动窗口即为s中涵盖了t的一个子字符串。
  validNum其实就是个快捷方式，否则就需要遍历t中的每一个字符，在滑动窗口中逐个比对，判断是否已经涵盖t中所有字符。
- ri：最小覆盖子串的起始索引；
- rl：最小覆盖子串的长度。初始化为题目中给出的max(s.length, t.length) + 1，即100_001。

处理过程：
维护滑动窗口：
- 不断扩展右边界：
   1. 若右侧元素在t中出现，则统计入windowCnt。若此时该字符串【在windowCnt中的出现次数 == 在targetCnt中的出现次数】，则validNum自增；
   2. 扩展右边界：for j := 0; j < len(nums); j++
- 当滑动窗口涵盖t，即validNum == len(targetCnt)时，试图在该下尽量使得s子串长度最小：
   1. 不断更新ri和rl的结果；
   2. 将左侧元素排除出滑动窗口；
	   - 因为左侧元素不一定在t中出现，即不一定加入windowCnt，所以需要做一道前置判断；
	   - 如果在windowCnt中，则需要判断其离开滑动窗口后，validNum是否需要自减；
	   - 最后再将其排除出windowCnt。注意上一步和该步有严格的顺序要求；
   3. 缩小左边界：i++
可以看出，扩展&缩小是一对对称操作。
*/

/**
- 字符串的处理：
	- 遍历string中的每一位：
		- for i：字节切片，可能无法正确处理非ASCII码字符；
 		- for range：Unicode码点。
	- 转换为string：string(a)
*/

/**
- 移动right扩大窗口时，应更新哪些数据？
  增加window计数器；
  更新valid；
- 何时应该暂停扩大窗口，转而缩小窗口？
  当t中所有字符被覆盖，即得到可行的覆盖子串时，需要收缩窗口以得到最小覆盖子串。
- 移动left缩小窗口时，应更新哪些数据？
  更新valid；
  减少window计数器；
  这两步在扩大窗口 & 收缩窗口时是完全对称的。
- 结果收集应在扩大窗口时还是缩小窗口时？
  在收缩窗口阶段更新最小覆盖子串，因为此时滑动窗口内的字符串是可行解，可以从中选取最优解。
*/
