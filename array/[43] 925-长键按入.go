package array

// 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
//
// 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
//
// 示例 1：
//
// 输入：name = "alex", typed = "aaleex"
// 输出：true
// 解释：'alex' 中的 'a' 和 'e' 被长按。
//
// 示例 2：
//
// 输入：name = "saeed", typed = "ssaaedd"
// 输出：false
// 解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
//
// 提示：
//
// 1 <= name.length, typed.length <= 1000
// name 和 typed 的字符都是小写字母
func isLongPressedName(name string, typed string) bool {
	ln, lt := len(name), len(typed)
	if ln > lt {
		return false
	}
	ni := 0
	var lastMatch byte
	for ti := 0; ti < lt; ti++ {
		if ni < ln && name[ni] == typed[ti] {
			lastMatch = name[ni]
			ni++
		} else if typed[ti] != lastMatch {
			return false
		}
	}
	return ni == len(name)
}

// 隐含 else if typed[ti] == lastMatch { continue }

/**
简单做法：
遍历typed -> 入栈消除重复元素 -> 比较name和栈最终残存字符串是否相等

双指针做法：
遍历typed，不断推进typed指针
	-> 若匹配，name指针推进，记录当前匹配的name中的字符至lastMatch
	-> 若不匹配，对比typed当前字符和lastMatch
		-> 若相等，则为"长键按入"的重复项，即else if typed[ti] == lastMatch { continue }
		-> 若不相等，则直接返回false
最后，判断name指针是否已扫遍name，
否则对于name = nb，typed = nnn无法通过。

需注意的细节：
if ni < ln && name[ni] == typed[ti]
ni < ln重要。若typed是长键按入，则len(name) <= len(typed)，且name一定比typed先遍历完，
在typed遍历过程中，需判断name已到达末尾，防止数组越界。

类似的双指针去重题，最好用一个变量做记录，简化处理。
一开始的思路是，如果name和typed字符串匹配，
则先推进name和typed指针，再不断推动typed指针来跳过重复字符，这样会引入一些边界条件，不好处理。
*/
