package sort

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var flag bool
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
