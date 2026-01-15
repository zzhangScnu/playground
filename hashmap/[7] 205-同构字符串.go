package hashmap

// 给定两个字符串 s 和 t ，判断它们是否是同构的。
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
//
// 示例 1:
//
// 输入：s = "egg", t = "add"
// 输出：true
//
// 示例 2：
//
// 输入：s = "foo", t = "bar"
// 输出：false
//
// 示例 3：
//
// 输入：s = "paper", t = "title"
// 输出：true
//
// 提示：
//
// 1 <= s.length <= 5 * 10⁴
// t.length == s.length
// s 和 t 由任意有效的 ASCII 字符组成
func isIsomorphic(s string, t string) bool {
	ms, mt := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		si, ti := s[i], t[i]
		if psi, ok := ms[si]; ok && psi != ti {
			return false
		}
		if pti, ok := mt[ti]; ok && pti != si {
			return false
		}
		ms[si] = ti
		mt[ti] = si
	}
	return true
}

/**
明确题意：
每个出现的字符都应当映射到另一个字符
	-> map1：s -> t，正向映射：保证每一个s字符唯一映射到t字符

不同字符不能映射到同一个字符上
	对于map1来说，无法在O(1)的时间复杂度内判断重复，所以需要额外的map
	-> map2：t -> s，逆向映射：保证每一个t字符唯一映射到s字符

结合map1和map2，可保证s <-> t双向映射唯一，即每个s字符只能唯一映射到一个t字符。

思路：
一边遍历长度相同的s和t，一边判断映射唯一性，再将s字符和t字符写入各自的map中。

如果只建立单向映射的bad case：
当 s = "ab"，t = "aa" 时：
原函数执行流程：
遍历 s[0] = 'a'，映射中无该键，建立 'a'→'a'；
遍历 s[1] = 'b'，映射中无该键，建立 'b'→'a'；
函数返回 true，但实际上这两个字符串不是同构的（t 的 'a' 同时对应 s 的 'a' 和 'b'）。
*/
