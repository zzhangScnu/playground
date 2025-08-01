package array

import (
	"math"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	res := math.MinInt
	h, count := citations[0], 1
	for i := 1; i < len(citations); i++ {
		h = citations[i]
		count++
		if h == count {
			res = max(res, h)
		}
	}
	return res
}
