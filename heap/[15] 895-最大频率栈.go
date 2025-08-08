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

/**
思路：
最大频率栈，虽然名称中带有"栈"，但其实不含先进后出特性，只弹出出现频率最高的元素。
如果频率相同，则弹出"最接近栈顶"的元素，即最新加入的元素。

需要维护：
KF：元素 -> 出现频率；
FK：出现频率 -> 元素，元素间通过链表连接，先进后出；
MaxFrequency：最大频率。

当增加元素时：
1. 需同时修改以上数据结构；
2. 当元素频率增加时，理论上需将元素从FK中移动。但举例：
元素m当前出现频率为1 -> FK [1 -> m]
	-> 压入元素m，出现频率为2 -> FK [1 -> m, 2 -> m]
	-> 弹出最大频率元素，此时是频率2对应的m -> FK [1 -> m]
	-> 弹出最大频率元素，此时是频率1对应的m -> FK []
只有这样，才能正确地将每个最大频率对应的元素弹出。

当弹出元素时：
1. 通过MaxFrequency + FK，定位到元素列表，从末尾取最新加入的元素返回；
2. 需同时修改以上数据结构。
*/

// 与155思想类似，需保留每一个位置的特定值。
