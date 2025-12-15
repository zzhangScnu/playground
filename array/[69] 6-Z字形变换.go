package array

import "strings"

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
// 请你实现这个将字符串进行指定行数变换的函数：
//
// string convert(string s, int numRows);
//
// 示例 1：
//
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
//
// 示例 2：
//
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// 示例 3：
//
// 输入：s = "A", numRows = 1
// 输出："A"
//
// 提示：
//
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	direction := -1
	row := 0
	for _, ch := range s {
		rows[row].WriteRune(ch)
		if row == 0 || row == numRows-1 {
			direction *= -1
		}
		row += direction
	}
	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}
	return res.String()
}

/**
思路：模拟
- 用 direction 控制方向（-1 向上 / 1 向下）
- 用 row 控制当前放置字符的行

处理过程：遍历所有字符：
- 当 direction 为 1 时，表示向下，此时 row += direction，row 递增，控制不断向下放置字符；
- 当 direction 为 -1 时，表示向上，此时 row += direction，row 递减，控制不断向上放置字符。
*/
