package diffarray

/**
假设你有一个长度为n的数组，初始情况下所有的数字匀为0，你将会被给出k个更新的操作。
其中，每个操作会被表示为一个三元组：[startIndex，endIndex，inc]，你需要将子数组
A[startIndex... endIndex] (包括 startIndex和 endIndex)增加inc。
请你返回k次操作后的数组。

示例1：
输入： length = 5， updates = [1，3，2]，[2，4，3]，[0，2，-2]1
输出：[-2，0，3，5，3]
解释：
初始状态：
[0，0，0，0，0]
进行了操作[1，3，2]后的状态：
[0，2，2，2，0]
进行了操作[2，4，3]后的状态：
[0，2，5，5，3]
进行了操作[0，2，-2]后的状态：
[-2，0，3，5，3]
*/

type Difference struct {
	diff []int
}

func NewDifference(n int) *Difference {
	return &Difference{
		diff: make([]int, n+1),
	}
}

func (d *Difference) Update(i, j, val int) {
	if i < 0 || j > len(d.diff)-1 || i > j {
		return
	}
	d.diff[i] += val
	d.diff[j+1] -= val
}

func (d *Difference) GetResult() []int {
	res := make([]int, len(d.diff)-1)
	res[0] = d.diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
