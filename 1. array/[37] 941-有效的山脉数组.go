package array

// 给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。
//
// 让我们回顾一下，如果 arr 满足下述条件，那么它是一个山脉数组：
//
// arr.length >= 3
// 在 0 < i < arr.length - 1 条件下，存在 i 使得：
//
// arr[0] < arr[1] < ... arr[i-1] < arr[i]
// arr[i] > arr[i+1] > ... > arr[arr.length - 1]
//
// 示例 1：
//
// 输入：arr = [2,1]
// 输出：false
//
// 示例 2：
//
// 输入：arr = [3,5,5]
// 输出：false
//
// 示例 3：
//
// 输入：arr = [0,3,2,1]
// 输出：true
//
// 提示：
//
// 1 <= arr.length <= 10⁴
// 0 <= arr[i] <= 10⁴
func validMountainArray(arr []int) bool {
	n := len(arr)
	highest := -1
	for i := 1; i < n && arr[i-1] < arr[i]; i++ {
		highest = i
	}
	if highest == -1 || highest == n-1 {
		return false
	}
	for i := highest + 1; i < n; i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}
	return true
} // todo: 简单的上升下降

func validMountainArrayTwoPointer(arr []int) bool {
	n := len(arr)
	left, right := 0, n-1
	for left < n-1 && arr[left] < arr[left+1] {
		left++
	}
	for right > 0 && arr[right-1] > arr[right] {
		right--
	}
	return left == right && left != 0 && right != n-1
} // todo：   return (left-1 == right+1) && left != 1 && right != n-2
