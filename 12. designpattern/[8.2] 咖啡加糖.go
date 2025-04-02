package designpattern

import "fmt"

/**
题目描述
小明喜欢品尝不同口味的咖啡，他发现每种咖啡都可以加入不同的调料，比如牛奶、糖和巧克力。他决定使用装饰者模式制作自己喜欢的咖啡。
请设计一个简单的咖啡制作系统，使用装饰者模式为咖啡添加不同的调料。系统支持两种咖啡类型：黑咖啡（Black Coffee）和拿铁（Latte）。

输入描述
多行输入，每行包含两个数字。第一个数字表示咖啡的选择（1 表示黑咖啡，2 表示拿铁），第二个数字表示要添加的调料类型（1 表示牛奶，2 表示糖）。

输出描述
根据每行输入，输出制作咖啡的过程，包括咖啡类型和添加的调料。

输入示例
1 1
2 2

输出示例
Brewing Black Coffee
Adding Milk
Brewing Latte
Adding Sugar
*/

type Coffee interface {
	Brew()
}

type BlackCoffee struct {
}

func (b BlackCoffee) Brew() {
	fmt.Println("Brewing Black Coffee")
}

type LatteCoffee struct {
}

func (l LatteCoffee) Brew() {
	fmt.Println("Brewing Latte")
}

type CondimentDecorator interface {
	Brew()
}

type MilkDecorator struct {
	Coffee Coffee
}

func (m MilkDecorator) Brew() {
	m.Coffee.Brew()
	fmt.Println("Adding Milk")
}

type SugarDecorator struct {
	Coffee Coffee
}

func (s SugarDecorator) Brew() {
	s.Coffee.Brew()
	fmt.Println("Adding Sugar")
}

func main() {
	for {
		var coffeeType, condimentType int
		fmt.Scan(&coffeeType, &condimentType)
		if coffeeType == 0 || condimentType == 0 {
			return
		}
		var coffee Coffee
		var condimentDecorator CondimentDecorator
		switch coffeeType {
		case 1:
			coffee = BlackCoffee{}
		case 2:
			coffee = LatteCoffee{}
		}
		switch condimentType {
		case 1:
			condimentDecorator = MilkDecorator{Coffee: coffee}
		case 2:
			condimentDecorator = SugarDecorator{Coffee: coffee}
		}
		condimentDecorator.Brew()
	}
}
