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
}

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
}

/**
思路一：for循环从左到右
其实可以在validMountainArray的基础上精简，仅使用一个游标，指定循环条件，遍历数组。
如果循环结束时，游标到达尾部，说明满足上升&下降趋势。

思路二：双指针从外向内
如果左右指针在中间相遇，说明左指针走过的路径均为上升，右指针走过的路径均为下降。
注意点：
1. 如果左指针在最左，说明全为下降；右指针在最右，说明全为上升；
2. 如果left = 0，left跟left+1比较，则left最终会停留在至高点；
   同理，right = n-1，right跟right-1比较，则right最终会停留在至高点。
   所以判断条件是left == right；
   如果left = 1，left跟left-1比较，则left最终会停留在至高点+1，因为满足上升的最后一次循环中，left会+1；
   同理，如果right = n-2，right跟right+1比较，则right最终会停留在至高点-1，因为满足下降的最后一次循环中，right会-1；
   所以判断条件是left - 1 == right + 1。
*/
