package stack_queue

import "unicode"

// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 整数除法仅保留整数部分。
//
// 你可以假设给定的表达式总是有效的。所有中间结果将在 [-2³¹, 2³¹ - 1] 的范围内。
//
// 注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
// 示例 1：
//
// 输入：s = "3+2*2"
// 输出：7
//
// 示例 2：
//
// 输入：s = " 3/2 "
// 输出：1
//
// 示例 3：
//
// 输入：s = " 3+5 / 2 "
// 输出：5
//
// 提示：
//
// 1 <= s.length <= 3 * 10⁵
// s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
// s 表示一个 有效表达式
// 表达式中的所有整数都是非负整数，且在范围 [0, 2³¹ - 1] 内
// 题目数据保证答案是一个 32-bit 整数
func calculateII(s string) int {
	var pre, cur, res int
	operator := '+'
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			i--
			if operator == '+' {
				res += cur
				pre = cur
			} else if operator == '-' {
				res -= cur
				pre = -cur
			} else if operator == '*' {
				res -= pre
				res += pre * cur
				pre = pre * cur
			} else if operator == '/' {
				res -= pre
				res += pre / cur
				pre = pre / cur
			}
			cur = 0
		} else if s[i] != ' ' {
			operator = int32(s[i])
		}
		i++
	}
	return res
}

/**
变量作用：
pre：
保存前一个操作数，记录【上一个需要参与乘除运算的数值】。
做乘法操作时，可能前一个操作数已经参与到了前置的加法/减法等较低优先级的运算中。
此时pre保存现场，以便修正先前的加减运算，高优处理当前的乘除运算。
举例说明：10 + 2 * 3
- 解析到 2 时，cur = 2，res = 10 + 2 = 12，此时令 pre = 2；
- 遇到 * 后解析到 3 时，需要用 pre 和当前数，即 2 和 3 做乘法，再修正 res，令 res = 12 - 2 + 2 * 3 = 14。

cur：
当前遍历到的操作数

res：
整体计算结果
*/

/**
注意：
循环获取任意位数的数字：
if unicode.IsDigit(rune(s[i])) {
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		cur = cur*10 + int(s[i]-'0') // 随着位数向右推进，累加当前数字。注意不是+=
		i++
	}
	i-- // 在最后一次循环中，i也会执行++，此时会指向数字的下一个运算符，所以这里需要回退一位
}
*/
