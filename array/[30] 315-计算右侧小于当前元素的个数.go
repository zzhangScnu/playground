package array

// 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
// nums[i] 的元素的数量。
//
// 示例 1：
//
// 输入：nums = [5,2,6,1]
// 输出：[2,1,1,0]
// 解释：
// 5 的右侧有 2 个更小的元素 (2 和 1)
// 2 的右侧仅有 1 个更小的元素 (1)
// 6 的右侧有 1 个更小的元素 (1)
// 1 的右侧有 0 个更小的元素
//
// 示例 2：
//
// 输入：nums = [-1]
// 输出：[0]
//
// 示例 3：
//
// 输入：nums = [-1,-1]
// 输出：[0,0]
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	var sorted []int
	for i := len(nums) - 1; i >= 0; i-- {
		index := findFarLeftIndex(sorted, nums[i])
		sorted = insert(sorted, index, nums[i])
		res[i] = index
	}
	return res
}

func findFarLeftIndex(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func insert(nums []int, index int, element int) []int {
	if len(nums) == 0 {
		return []int{element}
	}
	nums = append(nums, 0)
	copy(nums[index+1:], nums[index:])
	nums[index] = element
	return nums
}

/**
思路：
维护一个新的非递减序列sorted。
从后往前，对原数组中的每个元素i做二分查找，找到其在sorted中的位置index。
则sorted中[0, index)均为i右侧小于i的元素。
需要注意，因为要找i右侧【小于】i的元素，而不是小于等于i的元素，所以应该i在sorted中的插入位置应该是在最左侧，
如i = 3'，sorted = [1, 2, 3, 3, 3]，则插入后sorted = [1, 2, 3', 3, 3, 3]。

至于为什么是从后往前遍历原数组：
因为需要找i【右侧】小于i的元素个数。
倒序遍历，保证在遍历到i时，[i+1, len(nums)-1]的元素均在sorted有序，则i插入sorted的位置是相对【其在原数组的后续元素】有序的，
此时可以计算出原数组中有多少元素比它小。

扩展一下可得：
从后往前，将原数组非递减排序，i在sorted中的插入位置在最左侧-> [0, index)为i右侧小于i的元素；
从后往前，将原数组非递增排序，i在sorted中的插入位置在最右侧-> (index, len(nums)-1]为i右侧小于i的元素；

从前往后，将原数组非递减排序，i在sorted中的插入位置在最左侧-> [0, index)为i左侧小于i的元素；
从前往后，将原数组非递增排序，i在sorted中的插入位置在最右侧-> (index, len(nums)-1]为i左侧小于i的元素。

小技巧：
初始化res := make([]int, len(nums))，倒着填入结果，
而不是正着append，最后再reverse。

空出某个位置且插入某个数据：
1. 先对原数组append以增加其长度；
2. 对某位置后的数据整体向后移动一位；
3. 对某位置直接赋值；
4. 返回操作后的数据。
注意，返回值重要！

因为Go只有值传递，传入函数的是slice结构体的副本。
其中包含切片的长度、容量、指向底层数组的指针等，这些都是拷贝后的值。
但因为拷贝的指针仍与原指针指向同一个底层数组，通过其对底层数组的修改是可以生效且影响到原指针的访问的。
但有一种情况例外！就是append且操作后长度超出了容量，需要扩容的情况。
这时候会分配一个新的底层数组来承接元素，拷贝的指针也会指向这个新的底层数组。而原指针仍指向原底层数组。
两者已经不一样了，函数内的修改无法影响函数外。
*/
