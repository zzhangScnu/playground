package designpattern

import "fmt"

/**
题目描述
小明家的超市推出了不同的购物优惠策略，你可以根据自己的需求选择不同的优惠方式。其中，有两种主要的优惠策略：
1. 九折优惠策略：原价的90%。
2. 满减优惠策略：购物满一定金额时，可以享受相应的减免优惠。
具体的满减规则如下：
满100元减5元
满150元减15元
满200元减25元
满300元减40元
请你设计一个购物优惠系统，用户输入商品的原价和选择的优惠策略编号，系统输出计算后的价格。

输入描述
输入的第一行是一个整数 N（1 ≤ N ≤ 20），表示需要计算优惠的次数。
接下来的 N 行，每行输入两个整数，第一个整数M( 0 < M < 400) 表示商品的价格, 第二个整数表示优惠策略，1表示九折优惠策略，2表示满减优惠策略

输出描述
每行输出一个数字，表示优惠后商品的价格

输入示例
4
100 1
200 2
300 1
300 2

输出示例
90
175
270
260
*/

type Strategy interface {
	Do(cost int)
}

type DiscountStrategy struct {
}

func (d DiscountStrategy) Do(cost int) {
	fmt.Println(float64(cost) * 0.9)
}

type SpendAndSaveStrategy struct {
	thresholds []int
	save       []int
}

func (s SpendAndSaveStrategy) Do(cost int) {
	n := len(s.thresholds)
	for i := n - 1; i >= 0; i-- {
		if cost >= s.thresholds[i] {
			fmt.Println(cost - s.save[i])
			return
		}
	}
	fmt.Println(cost)
}

type Context struct {
	strategies map[int]Strategy
}

func NewContext() Context {
	return Context{
		strategies: map[int]Strategy{
			1: DiscountStrategy{},
			2: SpendAndSaveStrategy{
				thresholds: []int{100, 150, 200, 300},
				save:       []int{5, 15, 25, 40},
			},
		},
	}
}

func (c Context) Do(strategyType, cost int) {
	strategy := c.strategies[strategyType]
	if strategy == nil {
		fmt.Println("Unknown strategy")
	}
	strategy.Do(cost)
}

func main() {
	context := NewContext()
	var count int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var cost, strategyType int
		fmt.Scan(&cost, &strategyType)
		context.Do(strategyType, cost)
	}
}

/**
原本的实现是hardcode：
func (s SpendAndSaveStrategy) Do(cost int) {
	if cost >= 100 && cost < 150 {
		cost -= 5
	} else if cost >= 150 && cost < 200 {
		cost -= 15
	} else if cost >= 200 && cost < 300 {
		cost -= 25
	} else if cost >= 300 {
		cost -= 40
	}
	fmt.Println(cost)
}

可以改为用2个数组巧妙实现。
注意需从后往前遍历，使用贪心思想，先满足慢减金额大的情况，再向下降级匹配。
*/
