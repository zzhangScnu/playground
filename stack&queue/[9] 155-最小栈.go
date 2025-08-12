package stack_queue

//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
//
//
//
// 示例 1:
//
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// -2³¹ <= val <= 2³¹ - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 10⁴ 次

type MinStack struct {
	values           []int
	currentMinValues []int
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	currentMinValue := val
	if len(this.currentMinValues) > 0 && val > this.currentMinValues[len(this.currentMinValues)-1] {
		currentMinValue = this.currentMinValues[len(this.currentMinValues)-1]
	}
	this.currentMinValues = append(this.currentMinValues, currentMinValue)
}

func (this *MinStack) Pop() {
	this.values = this.values[0 : len(this.values)-1]
	this.currentMinValues = this.currentMinValues[0 : len(this.currentMinValues)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.currentMinValues[len(this.currentMinValues)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/**
思路：
引入辅助栈，用于记录新元素入栈时，栈中的最小元素。
意味着当nums[i]入栈时，values[i]存放nums[i]，而currentMinValues存放包含nums[i]在内的栈中所有元素的最小值。
该最小值可以通过与每个元素比较，取较小值进行维护。
这个比较操作需注意数组越界情况。
一开始我是这么写的：
func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	// 下面这行当栈为空时会导致数组越界
	currentMinValue := this.currentMinValues[len(this.currentMinValues)-1]
	if len(this.currentMinValues) == 0 || val < currentMinValue {
		currentMinValue = val
	}
	this.currentMinValues = append(this.currentMinValues, currentMinValue)
}
*/
