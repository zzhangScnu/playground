package slidingwindow

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
// 示例 1：
//
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
//
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
//
// 提示：
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
func checkInclusion(s1 string, s2 string) bool {
	targetCnt, windowCnt := make(map[string]int), make(map[string]int)
	for _, ch := range s1 {
		targetCnt[string(ch)]++
	}
	var i, valid int
	for j, ch := range s2 {
		rc := string(ch)
		if targetCnt[rc] > 0 {
			windowCnt[rc]++
			if windowCnt[rc] == targetCnt[rc] {
				valid++
			}
		}
		for len(targetCnt) == valid {
			if j-i+1 == len(s1) {
				return true
			}
			lc := string(s2[i])
			if targetCnt[lc] > 0 {
				if windowCnt[lc] == targetCnt[lc] {
					valid--
				}
				windowCnt[lc]--
			}
			i++
		}
	}
	return false
}

/**
思路：
- 滑动窗口：在s2中截取且不断从左向右滑动的子字符串；
- i & j：s2中滑动窗口的左 / 右边界索引；
- targetCnt：字符串s1中的每个字符 -> 出现次数；
- windowCnt：s2滑动窗口中的每个字符 -> 出现次数；
- valid：s2滑动窗口中符合s1中出现要求的字符的个数。
  当valid == len(targetCnt)时，表示s1中所有字符均在s2滑动窗口中涵盖，且出现次数与s1中一致，
  此时的滑动窗口即为s2中涵盖了s1的一个子字符串。
  valid其实就是个快捷方式，否则就需要遍历s1中的每一个字符，在s2滑动窗口中逐个比对，判断是否已经涵盖s1中所有字符。

处理过程：
维护滑动窗口：
- 不断扩展右边界：
   1. 若右侧元素在s1中出现，则统计入windowCnt。若此时该字符串【在windowCnt中的出现次数 == 在targetCnt中的出现次数】，则valid自增；
   2. 扩展右边界：for j, ch := range s2
- 当s2滑动窗口涵盖s1，即valid == len(targetCnt)时，更进一步判断s2的当前滑动窗口子串是否s1的一种排列：
   1. 若j-i+1 == len(s1)，则直接返回true；
   2. 否则将左侧元素排除出滑动窗口；
	   - 因为左侧元素不一定在t中出现，即不一定加入windowCnt，所以需要做一道前置判断；
	   - 如果在windowCnt中，则需要判断其离开滑动窗口后，valid是否需要自减；
	   - 最后再将其排除出windowCnt。注意上一步和该步有严格的顺序要求；
   3. 缩小左边界：i++
可以看出，扩展&缩小是一对对称操作。

可以看出，本题跟76的核心区别就是，如何判断满足题目中的约束条件：
s2窗口中涵盖s1所有字符 && s2窗口长度等于s1。
*/

/**
- 移动right扩大窗口时，应更新哪些数据？
  增加window计数器；
  更新valid；
- 何时应该暂停扩大窗口，转而缩小窗口？
  当t中所有字符被覆盖，即得到可行的覆盖子串时，需要收缩窗口来进一步判断是否存在长度相等的排列。
- 移动left缩小窗口时，应更新哪些数据？
  更新valid；
  减少window计数器；
  这两步在扩大窗口 & 收缩窗口时是完全对称的。
- 结果收集应在扩大窗口时还是缩小窗口时？
  在收缩窗口阶段判断s2是否包含s1的排列，因为此时滑动窗口内的字符串是可行解，可以从中选取最优解。
*/
