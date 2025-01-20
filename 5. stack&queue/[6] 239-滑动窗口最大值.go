package stack_queue

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
// 。
//
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
func maxSlidingWindow(nums []int, k int) []int {
	queue := DescendingQueue{
		data:     make([]int, 0, k),
		capacity: k,
	}
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	var res []int
	for i := k; i < len(nums); i++ {
		res = append(res, queue.GetMaxVal())
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
	}
	return res
}

type DescendingQueue struct {
	data     []int
	capacity int
}

func (d DescendingQueue) GetMaxVal() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

func (d DescendingQueue) Pop(num int) {
	if len(d.data) > 0 && d.data[0] == num {
		d.data = d.data[1:]
	}
}

func (d DescendingQueue) Push(num int) {
	if len(d.data) < d.capacity {
		d.data = append(d.data, num)
		return
	}
	idx := len(d.data) - 1
	for d.data[idx] <= num {
		d.data = d.data[:idx]
		idx--
	}
	d.data = append(d.data, num)
}
