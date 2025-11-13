package stack_queue

import "math"

// 给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。
//
// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
// 找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远
// 不会发生碰撞。
//
// 示例 1：
//
// 输入：asteroids = [5,10,-5]
// 输出：[5,10]
// 解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
//
// 示例 2：
//
// 输入：asteroids = [8,-8]
// 输出：[]
// 解释：8 和 -8 碰撞后，两者都发生爆炸。
//
// 示例 3：
//
// 输入：asteroids = [10,2,-5]
// 输出：[10]
// 解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
//
// 示例 4：
//
// 输入：asteroids = [3,5,-6,2,-1,4]
// 输出：[-6,2,4]
// 解释：小行星 -6 使小行星 3 和 5 爆炸，然后继续向左移动。在另一边，小行星 2 使小行星 -1 爆炸，然后继续向右移动，没有碰撞小行星 4。
//
// 提示：
//
// 2 <= asteroids.length <= 10⁴
// -1000 <= asteroids[i] <= 1000
// asteroids[i] != 0
func asteroidCollision(asteroids []int) []int {
	stack := NewStack()
	for _, asteroid := range asteroids {
		survive := true
		for survive && asteroid < 0 && !stack.IsEmpty() && stack.Peek() > 0 {
			cur := stack.Peek()
			if math.Abs(float64(cur)) < math.Abs(float64(asteroid)) {
				stack.Pop()
			} else if math.Abs(float64(cur)) == math.Abs(float64(asteroid)) {
				stack.Pop()
				survive = false
			} else {
				survive = false
			}
		}
		if survive {
			stack.Push(asteroid)
		}
	}
	return stack.data
}

/**
思路：
- 只有相向运行的小行星才会碰撞，反向或同向的不会
- 碰撞时，绝对值较小的小行星会爆炸
	- 如果是栈口的小行星爆炸：需要将该行星出栈，且在栈中连锁查找所有可能继续被摧毁的小行星进行处理
	- 如果是待入栈的小行星爆炸：循环结束
	- 如果同时爆炸：需要将该行星出栈
	- 如果到最后待入栈的小行星还存活，则将其入栈
*/
