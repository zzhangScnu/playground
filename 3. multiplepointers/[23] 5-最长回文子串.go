package multiplepointers

// 给定一个包含大写字母和小写字母的字符串
// s ，返回 通过这些字母构造成的 最长的 回文串 的长度。
//
// 在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。
//
// 示例 1:
//
// 输入:s = "abccccdd"
// 输出:7
// 解释:
// 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// 示例 2:
//
// 输入:s = "a"
// 输出:1
// 解释：可以构造的最长回文串是"a"，它的长度是 1。
//
// 提示:
//
// 1 <= s.length <= 2000
// s 只由小写 和/或 大写英文字母组成
func longestPalindrome(s string) string {
	var res string
	for i := range s {
		odd := palindrome(s, i, i)
		even := palindrome(s, i, i+1)
		if len(odd) > len(res) {
			res = odd
		}
		if len(even) > len(res) {
			res = even
		}
	}
	return res
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

/**
1. palindrome：寻找以left & right为中心，向两边扩展的回文子串
- 当left == right时，表示寻找奇数长度的子串；否则表示寻找偶数长度的子串；
- for里面指定左右指针遍历的合法索引范围，且结束时因为还有一次--和++，此时数组可能是越界的。
  而因为go的字符串切片操作是左闭右开的，所以截取答案时对left+1，而right不变。
2.longestPalindrome：遍历数组，分别寻找奇数、偶数长度的回文子串，且不断更新结果；
3. 【中心 -> 两端】的遍历方式少见，一般只有寻找回文中出现，其余大多是【两端 -> 中心】
*/
