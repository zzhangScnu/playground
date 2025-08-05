package array

import "strings"

// 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
//
// 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
// 注意:
//
// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。
//
// 示例 1:
//
// 输入: words = ["This", "is", "an", "example", "of", "text", "justification."],
// maxWidth = 16
// 输出:
// [
//
//	"This    is    an",
//	"example  of text",
//	"justification.  "
//
// ]
//
// 示例 2:
//
// 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// 输出:
// [
//
//	"What   must   be",
//	"acknowledgment  ",
//	"shall be        "
//
// ]
// 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//
//	因为最后一行应为左对齐，而不是左右两端对齐。
//	第二行同样为左对齐，这是因为这行只包含一个单词。
//
// 示例 3:
//
// 输入:words = ["Science","is","what","we","understand","well","enough","to",
// "explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 2
// 0
// 输出:
// [
//
//	"Science  is  what we",
//	"understand      well",
//	"enough to explain to",
//	"a  computer.  Art is",
//	"everything  else  we",
//	"do                  "
//
// ]
//
// 提示:
//
// 1 <= words.length <= 300
// 1 <= words[i].length <= 20
// words[i] 由小写英文字母和符号组成
// 1 <= maxWidth <= 100
// words[i].length <= maxWidth
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var line []string
	var length int
	for i := 0; i < len(words); i++ {
		if length+len(line)+len(words[i]) > maxWidth {
			spaceLength := maxWidth - length
			spaces := spaceLength / max(1, len(line)-1)
			extraSpaces := spaceLength % max(1, len(line)-1)
			for j := 0; j < max(1, len(line)-1); j++ {
				line[j] += strings.Repeat(" ", spaces)
				if extraSpaces > 0 {
					line[j] += " "
					extraSpaces--
				}
			}
			res = append(res, strings.Join(line, ""))
			line = make([]string, 0)
			length = 0
		}
		line = append(line, words[i])
		length += len(words[i])
	}
	lastLine := strings.Join(line, " ")
	spaceLength := maxWidth - len(lastLine)
	res = append(res, lastLine+strings.Repeat(" ", spaceLength))
	return res
}

/**
目标效果：
1. 对于前 n - 1 行，单词两端对齐，词间补位空格。若空格分配不均则从左往右开始补位；
2. 对于第 n 行，单词向左对齐，末尾补位空格。

思路：
- 使用贪心原则，在每一行中尽可能多地放置单词；
- 若无法继续放入下一个单词，则使用空格补齐本行；
	- 如果空格可以在单词间均匀分配，则直接插入单词之间；
	- 如果空格无法在单词间均匀分配，则将多余的空格从左往右插入单词之间；
- 最后再特殊处理最后一行，用空格补齐末尾部分。

代码分析：
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var line []string // 当前处理行，里面存放的是当前行的单词列表
	var length int // 当前处理行的长度，里面存放的是当前行的单词的长度和。注意此时尚未放置空格，故不包括空格长度
	for i := 0; i < len(words); i++ {
		// 若【当前行中放置的所有单词的长度和 + 当前行中放置的所有单词间的空格长度 + 下一个试图放入的单词长度 + 下一个试图放入的单词引入的空格】 > 行最大长度
		// 即【当前行中放置的所有单词的长度和 + (当前行中放置的所有单词的个数 - 1) + 下一个试图放入的单词长度 + 1】 > 行最大长度
		// 即【当前行中放置的所有单词的长度和 + 当前行中放置的所有单词的个数 + 下一个试图放入的单词长度】 > 行最大长度
		if length+len(line)+len(words[i]) > maxWidth {
			// 需要补齐的空格长度
			spaceLength := maxWidth - length
			// 均匀分配到每个单词间的空格长度
			// max(1, len(line)-1)，简洁兼容了每行中单词长度为1导致余数为0异常的场景
			spaces := spaceLength / max(1, len(line)-1)
			// 多出来的、需要从左到右分配到单词间的空格长度
			extraSpaces := spaceLength % max(1, len(line)-1)
			// 循环次数为[0, max(1, len(line)-1))，简洁兼容了每行中单词长度为1的场景，也能执行一次从而补齐空格
			for j := 0; j < max(1, len(line)-1); j++ {
				line[j] += strings.Repeat(" ", spaces)
				// 因为多出的空格是对max(1, len(line)-1)取模，最大循环次数也是max(1, len(line)-1)，所以能保证至少在循环结束前，多余空格能分配完
				if extraSpaces > 0 {
					line[j] += " "
					extraSpaces--
				}
			}
			res = append(res, strings.Join(line, ""))
			// 清零暂存区
			line = make([]string, 0)
			length = 0
		}
		// 贪心思想，如果本行能纳入下一个单词，则纳入
		line = append(line, words[i])
		length += len(words[i])
	}
	// 特殊处理最后一行
	// 先在单词间插入空格
	lastLine := strings.Join(line, " ")
	// 再在最后填充空格
	spaceLength := maxWidth - len(lastLine)
	res = append(res, lastLine+strings.Repeat(" ", spaceLength))
	return res
}
*/
