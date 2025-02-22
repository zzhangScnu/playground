package backtracking

import (
	"strconv"
	"strings"
)

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
// 和 "192.168@1.1" 是 无效 IP 地址。
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
// 排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
//
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
//
// 示例 3：
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
// 提示：
//
// 1 <= s.length <= 20
// s 仅由数字组成

var addresses []string

var addressesRes []string

func restoreIpAddresses(s string) []string {
	addresses, addressesRes = []string{}, []string{}
	doRestoreIpAddresses(s, 0)
	return addressesRes
}

func doRestoreIpAddresses(s string, beginIdx int) {
	if beginIdx >= len(s) && len(addresses) == 4 {
		addressesRes = append(addressesRes, strings.Join(addresses, "."))
		return
	}
	for i := beginIdx; i < len(s); i++ {
		if i != beginIdx && s[beginIdx] == '0' {
			continue
		}
		str := s[beginIdx : i+1]
		num, _ := strconv.Atoi(s[beginIdx : i+1])
		if num >= 0 && num <= 255 {
			addresses = append(addresses, str)
			doRestoreIpAddresses(s, i+1)
			addresses = addresses[:len(addresses)-1]
		}
	}
}

/**
跟131整体思路相似。
- 判断前导0：若切割出来的子字符串长度大于1且以0开始，则表示不是数字0但有前导0，不合法；
- base case：需要增加长度 == 4的限制条件。
*/
