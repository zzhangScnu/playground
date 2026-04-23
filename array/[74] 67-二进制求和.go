package array

import "strings"

// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
//
// 示例 1：
//
// 输入:a = "11", b = "1"
// 输出："100"
//
// 示例 2：
//
// 输入：a = "1010", b = "1011"
// 输出："10101"
//
// 提示：
//
// 1 <= a.length, b.length <= 10⁴
// a 和 b 仅由字符 '0' 或 '1' 组成
// 字符串如果不是 "0" ，就不含前导零
func addBinary(a string, b string) string {
	var sb strings.Builder
	i, j, carry := len(a)-1, len(b)-1, byte(0)
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += a[i] - '0'
			i--
		}
		if j >= 0 {
			carry += b[j] - '0'
			j--
		}
		sb.WriteByte(carry%2 + '0')
		carry /= 2
	}
	str := sb.String()
	bytes := []byte(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
