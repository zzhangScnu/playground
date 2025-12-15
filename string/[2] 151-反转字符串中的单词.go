package string

// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//
// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//
// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
//
// 示例 1：
//
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
//
// 示例 2：
//
// 输入：s = " hello world "
// 输出："world hello"
// 解释：反转后的字符串中不能存在前导空格和尾随空格。
//
// 示例 3：
//
// 输入：s = "a good  example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
//
// 提示：
//
// 1 <= s.length <= 10⁴
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词
//
// 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
func reverseWords(s string) string {
	s = removeExtraSpace(s)
	s = reverse(s, 0, len(s)-1)
	var beginIdx int
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			s = reverse(s, beginIdx, i-1)
			beginIdx = i + 1
		}
	}
	return s
}

func removeExtraSpace(s string) string {
	chars := []byte(s)
	var slow, fast int
	for ; fast < len(s); fast++ {
		if chars[fast] != ' ' {
			if slow != 0 {
				chars[slow] = ' '
				slow++
			}
			for fast < len(s) && chars[fast] != ' ' {
				chars[slow] = chars[fast]
				slow++
				fast++
			}
		}
	}
	chars = chars[:slow]
	return string(chars)
}

/**
整体思路：
1. 去掉多余空格-采用移除元素思路；
	- slow：指向新数组的末尾，即下一个元素插入的位置；
	- fast：指向符合条件的结果，即下一个元素；
	同时需要注意，每个单词之间的一个空格，是在移除所有空格后再次添加的。首单词前不需要添加，所以额外判断了一下slow的指向。
2. 反转整个字符串；
3. 反转每个单词；
	var beginIdx int
	for i := 0; i <= len(s); i++ { // i == len(s)时也要执行，否则最后一个单词无法反转，这里很容易漏掉
		if i == len(s) || s[i] == ' ' { // i == len(s)要放在前面，否则会导致数组越界
			s = reverse(s, beginIdx, i-1) // i此时是空格，反转区间应结束于i-1
			beginIdx = i + 1 // i此时是空格，应从i+1开始下一轮反转
		}
	}

字符串 -> 数组：[]byte(s)
数组 -> 字符串：string(chars)
切片：chars = chars[:slow]，且左闭右开
*/
