package stack_queue

import "unicode"

/*
*
Implement a basic calculator to evaluate a simple expression string
The expression string contains only non-negative integers, '+', '-', '*', '\' operators, and
open '(' and closing parentheses ')'. The integer division should truncate toward zero
You may assume that the given expression is always valid. AAll intermediate
results will be in the range of [-231, 231 - 1].
Note: You are not allowed to use any built-in function which evaluates strings
as mathematical expressions, such as eval().
Input: s = "2*(5+5*2)/3+(6/2+8)"
Output: 21
*/
func calculateIII(s string) int {
	pre, cur, res, operator := 0, 0, 0, 1
	var stack []int
	for i := 0; i < len(s); {
		ch := s[i]
		if unicode.IsDigit(rune(ch)) {
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				cur = cur*10 + int(s[i]-'0')
				i++
			}
			i--
			pre, cur, res = doCalculate(pre, cur, res, operator)
		} else {
			if ch != ' ' {
				switch ch {
				case '-':
					operator = -1
				case '+':
					operator = 1
				case '*':
					operator = 2
				case '/':
					operator = 3
				case '(':
					stack = append(stack, pre, res, operator)
					pre, cur, res, operator = 0, 0, 0, 1
				case ')':
					prevPre, prevRes, prevOperator := stack[len(stack)-3], stack[len(stack)-2], stack[len(stack)-1]
					stack = stack[:len(stack)-3]
					cur = res
					pre, res, operator = prevPre, prevRes, prevOperator
					pre, cur, res = doCalculate(pre, cur, res, operator)
				}
			}
			i++
		}
	}
	return res
}

func doCalculate(pre, cur, res, operator int) (int, int, int) {
	switch operator {
	case 1:
		res += cur
		pre = cur
		cur = 0
	case -1:
		res -= cur
		pre = -cur
		cur = 0
	case 2:
		res = res - pre + pre*cur
		pre = pre * cur
		cur = 0
	case 3:
		res = res - pre + pre/cur
		pre = pre / cur
		cur = 0
	}
	return pre, cur, res
}

// 在处理括号时将 `pre` 入栈，是为了正确恢复括号外的计算状态，确保括号内的结果能和括号外的前序运算正确衔接。具体原因如下：
//
//
//### 1. `pre` 的作用：记录「待参与乘除运算的中间结果」
//`pre` 变量的核心功能是保存「上一个需要参与乘除运算的数值」。例如：
//- 计算 `10 + 2 * 3` 时：
//  - 解析到 `2` 时，`pre = 2`，`res = 10 + 2 = 12`。
//  - 遇到 `*` 后解析到 `3` 时，需要用 `pre`（2）和当前数（3）做乘法，再修正 `res`（`res = 12 - 2 + 2*3 = 14`），此时 `pre` 更新为 `2*3=6`。
//- 可见，`pre` 是乘除运算的「桥梁」，记录着需要和下一个数进行乘除的中间值。
//
//
//### 2. 括号会打断运算流程，需要保存 `pre` 才能恢复
//当遇到左括号 `(` 时，括号内的运算会形成一个独立的子表达式（例如 `2 + 3*(4+5)` 中的 `4+5`）。此时：
//- 括号外的运算状态（`res`、`operator`、`pre`）需要被「暂停」，否则括号内的计算会覆盖这些变量。
//- 括号内的计算结束后，需要将结果作为一个「整体数值」，继续参与括号外的运算（例如 `3*(4+5)` 中，`4+5=9` 需作为整体和 `3` 相乘）。
//
//此时如果不保存 `pre`，会导致括号外的乘除运算无法正确衔接。例如：
//- 计算 `2 * (3 + 4)`：
//  - 遇到 `(` 时，括号外的状态是 `res=0`，`operator=2`（乘），`pre=2`（需要和括号内结果相乘的数）。
//  - 若不保存 `pre=2`，括号内计算完 `3+4=7` 后，无法知道需要用 `2 * 7` 得到最终结果。
//
//
//### 3. 栈的作用：完整保存「运算上下文」
//栈中保存 `res`、`operator`、`pre` 三个变量，本质是保存括号外的「完整运算上下文」：
//- `res`：括号外已累计的结果（例如 `10 + 2*(...)` 中的 `10`）。
//- `operator`：括号外等待执行的运算符（例如 `10 + 2*(...)` 中的 `+`）。
//- `pre`：括号外需要参与乘除的中间值（例如 `10 + 2*(...)` 中的 `2`）。
//
//当遇到右括号 `)` 时，从栈中恢复这三个变量，才能让括号内的结果（例如 `7`）正确参与括号外的运算（例如 `2 * 7`）。
//
//
//### 总结
//`pre` 是乘除运算的关键中间变量，记录着括号外「等待参与运算的数值」。如果不将其入栈，括号内的计算会覆盖该值，导致括号内外的运算无法正确衔接。因此，必须将 `pre` 与 `res`、`operator` 一起入栈，才能完整保存运算状态，确保括号处理的正确性。
