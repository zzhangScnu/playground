package monotonicstack

import . "code.byted.org/zhanglihua.river/playground/5. stack&queue"

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
// 在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
// 示例 1:
//
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出:[1,1,4,2,1,1,0,0]
//
// 示例 2:
//
// 输入: temperatures = [30,40,50,60]
// 输出:[1,1,1,0]
//
// 示例 3:
//
// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]
//
// 提示：
//
// 1 <= temperatures.length <= 10⁵
// 30 <= temperatures[i] <= 100
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	st := NewStack()
	st.Push(0)
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] <= st.Peek() {
			st.Push(i)
		} else {
			for !st.IsEmpty() && temperatures[i] > temperatures[st.Peek()] {
				index := st.Pop()
				res[index] = i - index
			}
			st.Push(i)
		}
	}
	return res
}

/**
思路：
需要求出每个元素右边第一个比其大的元素位置，可以使用单调栈。

递增/递减？
需要使用梯度递增的栈，才能找出元素右边第一个比其大的元素位置。

初始化？
需要保持默认零值，才能处理"后续没有比本元素更大的元素"的场景。

注意点：
1. 栈中存放的是索引，故比较时要重新取值；
2. for循环中务必检查栈非空。
*/
