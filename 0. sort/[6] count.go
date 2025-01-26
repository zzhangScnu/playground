package sort

import "math"

func countSort(nums []int) []int {
	minNum, maxNum := math.MaxInt, math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	rangeSize := maxNum - minNum + 1
	counts := make([]int, rangeSize)
	for _, num := range nums {
		counts[num-minNum]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] = counts[i-1] + counts[i]
	}
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		pos := counts[nums[i]-minNum] - 1
		res[pos] = nums[i]
		counts[nums[i]-minNum]--
	}
	return res
}

/**
计数排序适用于待排序数据集范围较固定且有一定排布规律的场景。
用到了前缀和数组的思想，将数据按大小放置于前后相对的位置，再按位置输出，无需比较而达到排序效果。
为什么最后是从后往前遍历数组来输出结果的呢？——为了保证排序的稳定性，原本靠后的元素，输出位置也靠后。
*/
