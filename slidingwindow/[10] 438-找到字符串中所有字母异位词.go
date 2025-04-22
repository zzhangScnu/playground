package slidingwindow

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 示例 1:
//
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
// 示例 2:
//
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
// 提示:
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
func findAnagrams(s string, p string) []int {
	windowCnt, targetCnt := make(map[string]int), make(map[string]int)
	for _, ch := range p {
		targetCnt[string(ch)]++
	}
	var i, valid int
	var res []int
	for j := 0; j < len(s); j++ {
		rc := string(s[j])
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				valid++
			}
		}
		for valid == len(targetCnt) {
			if j-i+1 == len(p) {
				res = append(res, i)
			}
			lc := string(s[i])
			if targetCnt[lc] > 0 {
				if windowCnt[lc] == targetCnt[lc] {
					valid--
				}
				windowCnt[lc]--
			}
			i++
		}
	}
	return res
}

/**
思路：
- 滑动窗口：在s中截取且不断从左向右滑动的子字符串；
- i & j：滑动窗口的左 / 右边界索引；
- targetCnt：字符串p中的每个字符 -> 出现次数；
- windowCnt：滑动窗口中的每个字符 -> 出现次数；
- valid：滑动窗口中符合t中出现要求的字符的个数。
  当valid == len(targetCnt)时，表示t中所有字符均在滑动窗口中涵盖，且出现次数与t中一致，
  此时的滑动窗口即为s中涵盖了p的一个子字符串。
  valid其实就是个快捷方式，否则就需要遍历t中的每一个字符，在滑动窗口中逐个比对，判断是否已经涵盖p中所有字符。

处理过程：
维护滑动窗口：
- 不断扩展右边界：
   1. 若右侧元素在t中出现，则统计入windowCnt。若此时该字符串【在windowCnt中的出现次数 == 在targetCnt中的出现次数】，则valid自增；
   2. 扩展右边界：for j := 0; j < len(s); j++
- 当滑动窗口涵盖t，即valid == len(targetCnt)时：
   1. 若此时滑动窗口长度j-i+1 == len(p)，表示当前滑动窗口截取的s的字符串与p是字母异位词，需要收集结果：res = append(res, i)
   2. 将左侧元素排除出滑动窗口；
	   - 因为左侧元素不一定在t中出现，即不一定加入windowCnt，所以需要做一道前置判断；
	   - 如果在windowCnt中，则需要判断其离开滑动窗口后，valid是否需要自减；
	   - 最后再将其排除出windowCnt。注意上一步和该步有严格的顺序要求；
   3. 缩小左边界：i++
可以看出，扩展&缩小是一对对称操作。

可以看出，本题跟567的核心区别就是，判断满足题目中的约束条件后，需要收集所有可能的结果：
res = append(res, i)
*/

/**
- 移动right扩大窗口时，应更新哪些数据？
  增加window计数器；
  更新valid；
- 何时应该暂停扩大窗口，转而缩小窗口？
  当t中所有字符被覆盖，即得到可行的覆盖子串时，需要收缩窗口以尝试寻找p的字母异位词。
- 移动left缩小窗口时，应更新哪些数据？
  更新valid；
  减少window计数器；
  这两步在扩大窗口 & 收缩窗口时是完全对称的。
- 结果收集应在扩大窗口时还是缩小窗口时？
  在收缩窗口阶段通过不断更新左边界，来得到可能的字母异位词。因为此时滑动窗口内的字符串是可行解，可以从中选取最优解。
*/
