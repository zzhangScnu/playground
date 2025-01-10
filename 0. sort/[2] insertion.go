package sort

func insertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ { // 元素
		var flag bool
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				swap(nums, j-1, j)
				flag = true
			}
			if !flag {
				break
			}
		}
	}
	return nums
}

func insertionSortII(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		value := nums[i]
		j := i - 1
		// 查找插⼊的位置
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j] // 数据移动
			} else {
				break
			}
		}
		nums[j+1] = value // 插⼊数据
	}
	return nums
}
