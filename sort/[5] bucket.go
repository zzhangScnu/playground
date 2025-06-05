package sort

import (
	"code.byted.org/zhanglihua.river/playground/sort/quick"
	"math"
)

func bucketSort(nums []int, bucketNum int) []int {
	buckets := make([][]int, bucketNum)
	minNum, maxNum := math.MaxInt, math.MinInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	bucketRange := (maxNum - minNum + 1) / bucketNum
	if bucketRange == 0 {
		bucketRange = 1
	}
	for _, num := range nums {
		bucket := (num - minNum + 1) / bucketRange
		if bucketRange >= bucketNum {
			bucket = bucketNum - 1
		}
		buckets[bucket] = append(buckets[bucket], num)
	}
	var res []int
	for _, bucket := range buckets {
		quick.quickSort(bucket)
		res = append(res, bucket...)
	}
	return res
}

/**
桶排序的思想：
将n个数尽可能均匀地分散在m个桶中，再对每个桶进行排序。
适用于待排序数据集范围较固定的场景，如按考试分数排序。

需要注意边界处理：
- 如果最大/最小值的差距很小，那么至少每个桶要装入一个元素，即桶的元素范围至少为1：
	if bucketRange == 0 {
		bucketRange = 1
	}
- 如果在分配元素时，算出的桶编号越界，则全部放入最大的桶中：
	if bucketRange >= bucketNum {
		bucket = bucketNum - 1
	}
	这里不能取模，否则会打乱元素间的相对顺序。
*/
