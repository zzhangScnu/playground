package heap

import "container/list"

//设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。
//
// 实现 FreqStack 类:
//
//
//
// FreqStack() 构造一个空的堆栈。
//
// void push(int val) 将一个整数 val 压入栈顶。
//
// int pop() 删除并返回堆栈中出现频率最高的元素。
//
// 如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。
//
//
//
//
//
// 示例 1：
//
//
//输入：
//["FreqStack","push","push","push","push","push","push","pop","pop","pop",
//"pop"],
//[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
//输出：[null,null,null,null,null,null,null,5,7,5,4]
//解释：
//FreqStack = new FreqStack();
//freqStack.push (5);//堆栈为 [5]
//freqStack.push (7);//堆栈是 [5,7]
//freqStack.push (5);//堆栈是 [5,7,5]
//freqStack.push (7);//堆栈是 [5,7,5,7]
//freqStack.push (4);//堆栈是 [5,7,5,7,4]
//freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
//freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
//freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
//freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
//freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。
//
//
//
// 提示：
//
//
// 0 <= val <= 10⁹
// push 和 pop 的操作数不大于 2 * 10⁴。
// 输入保证在调用 pop 之前堆栈中至少有一个元素。

type FreqStack struct {
	KF           map[int]int
	FK           map[int]*list.List
	MaxFrequency int
}

func FreqStackConstructor() FreqStack {
	return FreqStack{
		KF:           make(map[int]int),
		FK:           make(map[int]*list.List),
		MaxFrequency: 0,
	}
}

func (m *FreqStack) Push(val int) {
	m.KF[val]++
	if m.FK[m.KF[val]] == nil {
		m.FK[m.KF[val]] = list.New()
	}
	m.FK[m.KF[val]].PushBack(val)
	if m.MaxFrequency < m.KF[val] {
		m.MaxFrequency = m.KF[val]
	}
}

func (m *FreqStack) Pop() int {
	val := m.FK[m.MaxFrequency].Back().Value.(int)
	m.FK[m.MaxFrequency].Remove(m.FK[m.MaxFrequency].Back())
	m.KF[val]--
	if m.FK[m.MaxFrequency].Len() == 0 {
		m.MaxFrequency--
	}
	return val
}
