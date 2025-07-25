package divide_conquer

// 给定一个正整数 n，编写一个函数，获取一个正整数的二进制形式并返回其二进制表达式中 设置位 的个数（也被称为汉明重量）。
//
// 示例 1：
//
// 输入：n = 11
// 输出：3
// 解释：输入的二进制串 1011中，共有 3 个设置位。
//
// 示例 2：
//
// 输入：n = 128
// 输出：1
// 解释：输入的二进制串 10000000中，共有 1 个设置位。
//
// 示例 3：
//
// 输入：n = 2147483645
// 输出：30
// 解释：输入的二进制串 1111111111111111111111111111101 中，共有 30 个设置位。
//
// 提示：
//
// 1 <= n <= 2³¹ - 1
//
// 进阶：
//
// 如果多次调用这个函数，你将如何优化你的算法？
func hammingWeight(n int) int {
	var cnt int
	for n > 0 {
		n &= n - 1
		cnt++
	}
	return cnt
}

func hammingWeightDivideConquer(n int) int {
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0F0F0F0F) + ((n >> 4) & 0x0F0F0F0F)
	n = (n & 0x00FF00FF) + ((n >> 8) & 0x00FF00FF)
	n = (n & 0x0000FFFF) + (n >> 16)
	return n
}

/**
方法一：每次都消除n中末尾的1，并累加结果，直到n为0。循环次数与n中1的位数相等；
为什么能消除末尾的1？
如：
		   n  			0 1 0 0
	  n - 1  			0 0 1 1
n & (n - 1)				0 0 0 0 -> 将第2位即末尾的1消除


方法二：分治。
第一轮处理：统计每2位中1的数量：
- 统计奇数位中1的数量：(n & 0x55555555)
- 统计偶数位中1的数量-将偶数位移动到每2位中的低位，再与掩码与提取出来：((n >> 1) & 0x55555555)
- 两者相加，将结果保存在每2位中。
持续进行下轮处理，直到将高16位和低16位累加结果相加。
*/
