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
	res := []int{queue.GetMaxVal()}
	for i := k; i < len(nums); i++ {
		queue.Pop(nums[i-k])
		queue.Push(nums[i])
		res = append(res, queue.GetMaxVal())
	}
	return res
}

type DescendingQueue struct {
	data     []int
	capacity int
}

func (d *DescendingQueue) GetMaxVal() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

func (d *DescendingQueue) Pop(num int) {
	if len(d.data) > 0 && d.data[0] == num {
		d.data = d.data[1:]
	}
}

func (d *DescendingQueue) Push(num int) {
	for len(d.data) > 0 && d.data[len(d.data)-1] < num {
		d.data = d.data[:len(d.data)-1]
	}
	d.data = append(d.data, num)
}

/**
维护一个单调队列，维护一组单调非递减的数据。队列方法自定义：
1. GetMaxVal：获取队列中的最大值——通过窗口滑动+Pop+Push操作，队头即为最大元素；
2. Push：放入新值——队列只维护在窗口中“可能成为最大值”的元素，所以在进入窗口的元素入列时，
	要将前面所有比其【小】的元素都移除，因为它们在当前窗口内，不可能成为最大值；
3. Pop：将头节点弹出——如果头节点值等于离开窗口的元素值，则真实操作弹出。
	否则已经在之前被移除掉了，本次无需操作。
*/
