package array

// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而
// "aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代
// 码？
//
// 示例 1：
//
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
//
// 示例 2：
//
// 输入：s = "axc", t = "ahbgdc"
// 输出：false
//
// 提示：
//
// 0 <= s.length <= 100
// 0 <= t.length <= 10^4
// 两个字符串都只由小写字符组成。
func isSubsequenceBinarySearch(s string, t string) bool {
	indexes := make(map[int32][]int)
	for i, ch := range t {
		indexes[ch] = append(indexes[ch], i)
	}
	start := 0
	for _, ch := range s {
		low := searchFirstGE(ch, indexes, start)
		if low == -1 || low == len(indexes[ch]) {
			return false
		}
		start = indexes[ch][low] + 1
	}
	return true
}

func searchFirstGE(ch int32, indexes map[int32][]int, start int) int {
	targetIndexes := indexes[ch]
	if targetIndexes == nil {
		return -1
	}
	low, high := 0, len(targetIndexes)-1
	for low <= high {
		mid := low + (high-low)>>1
		if targetIndexes[mid] < start {
			low = mid + 1
		} else if targetIndexes[mid] >= start {
			high = mid - 1
		}
	}
	return low
}

/**
思路：
进阶版本：如果有大量输入的 S，称作 S1, S2, ... , Sk，其中 k >= 10亿，需要依次检查它们是否为 T 的子序列。
此时有大量的s，但只有一个t，意味着如果用双指针做法的时间复杂度会是O(M * N)。
但如果引入二分搜索，可以降为O(M * logN)。

核心是维护t中【字符 -> 索引列表】的映射关系，由于从前到后遍历t，索引列表天然有序。
维护一个变量start，指向t中的某个位置，表示t中字符与s中字符的当前最新匹配位置，的下一个位置。
即本轮从t的哪个位置开始与s尝试匹配。
对于s中的每一个字符ch，在t中ch -> 索引列表中进行二分搜索，目标是找到t中较start相等或更大的位置。此时：
- 如果s中该字符，在t中根本不存在，则直接返回false。即 low == -1 的情况；
- 如果遍历到索引列表末尾越界处，则表示没有找到，直接返回false。即 low == len(indexes[ch]) 的情况；
否则表示t中存在字符与s中的字符匹配成功，t中的匹配位置为indexes[ch][low]。
对其+1，从下个位置开始，对s的下一个字符进行新一轮匹配。
*/
