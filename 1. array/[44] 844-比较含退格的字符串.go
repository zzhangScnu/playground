package array

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。
//
// 注意：如果对空文本输入退格字符，文本继续为空。
//
// 示例 1：
//
// 输入：s = "ab#c", t = "ad#c"
// 输出：true
// 解释：s 和 t 都会变成 "ac"。
//
// 示例 2：
//
// 输入：s = "ab##", t = "c#d#"
// 输出：true
// 解释：s 和 t 都会变成 ""。
//
// 示例 3：
//
// 输入：s = "a#c", t = "b"
// 输出：false
// 解释：s 会变成 "c"，但 t 仍然是 "b"。
//
// 提示：
//
// 1 <= s.length, t.length <= 200
// s 和 t 只含有小写字母以及字符 '#'
//
// 进阶：
//
// 你可以用 O(n) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？
func backspaceCompare(s string, t string) bool {
	si, ti := len(s)-1, len(t)-1
	scnt, tcnt := 0, 0
	for si >= 0 || ti >= 0 {
		for si >= 0 {
			if s[si] == '#' {
				scnt++
				si--
			} else if scnt > 0 {
				scnt--
				si--
			} else {
				break
			}
		}
		for ti >= 0 {
			if t[ti] == '#' {
				tcnt++
				ti--
			} else if tcnt > 0 {
				tcnt--
				ti--
			} else {
				break
			}
		}
		if si < 0 || ti < 0 {
			break
		}
		if s[si] != t[ti] {
			return false
		}
		si--
		ti--
	}
	return si == -1 && ti == -1
}

/**
思路一：栈
【从前往后】顺序将字符入栈 -> 遇到"#"则弹出栈口元素，进行消除 -> 最终栈内即为退格后的字符串
分别处理s和t，再比较是否相等

思路二：双指针
【从后往前】遍历字符串
	-> 遇到"#"：count++；
	-> 遇到普通字符串：
		-> 若count > 0：移动指针来跳过字符，模拟消除
			-> 直到count == 0，此时指针指向的是当前处理完所有退格后的字符的位置
			-> 比较字符
				-> 若不相等：直接return false
				-> 若相等，推进指针，进行下一轮处理和比较
		-> 若count == 0：比较字符，逻辑同上
最后，如果两个指针都指向-1的位置，表示扫遍了s和t都是相等的，符合条件。

需注意的细节：
if si < 0 || ti < 0 {
	break
}
这个条件很重要，因为s和t可能会提前遍历完，要防止数组越界。
*/
