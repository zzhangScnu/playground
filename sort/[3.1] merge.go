package sort

var temp []int

func mergeSortII(nums []int) []int {
	n := len(nums)
	temp = make([]int, n)
	sort(nums, 0, n-1)
	return nums
}

func sort(nums []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)>>1
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	mergeII(nums, lo, mid, hi)
}

func mergeII(nums []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}
	i, j := lo, mid+1
	for cur := lo; cur <= hi; cur++ {
		if i == mid+1 {
			nums[cur] = temp[j]
			j++
		} else if j == hi+1 {
			nums[cur] = temp[i]
			i++
		} else if temp[i] < temp[j] {
			nums[cur] = temp[i]
			i++
		} else {
			nums[cur] = temp[j]
			j++
		}
	}
}
