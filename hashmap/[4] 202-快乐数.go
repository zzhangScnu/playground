package hashmap

// 编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
//
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
//
// 示例 1：
//
// 输入：n = 19
// 输出：true
// 解释：
// 1² + 9² = 82
// 8² + 2² = 68
// 6² + 8² = 100
// 1² + 0² + 0² = 1
//
// 示例 2：
//
// 输入：n = 2
// 输出：false
//
// 提示：
//
// 1 <= n <= 2³¹ - 1
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = cal(slow)
		fast = cal(fast)
		fast = cal(fast)
		if slow == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func cal(n int) int {
	var res int
	for n > 0 {
		bit := n % 10
		res += bit * bit
		n /= 10
	}
	return res
}

/**
解法1-哈希：当计算结果不为1时，用哈希记录。当重复出现结果时，即表示陷入无限循环，不是快乐数；
解法2-快慢指针：快指针走两步，慢指针走一步。若不为快乐数，则表示结果集有环，快慢指针一定会相遇（计算结果相同且不为1）。
*/
