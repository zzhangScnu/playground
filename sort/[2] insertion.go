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

/**
原地
稳定
最好O(n)，最坏O(n2)

为什么时间复杂度插入和冒泡相同，而插入排序性能更好？
因为每次交换，冒泡需要将元素移动3次；
而插入只需移动1次即可。
*/
