package designpattern

import "fmt"

/**
题目描述
小明去奶茶店买奶茶，他可以通过在自助点餐机上来点不同的饮品，请你使用命令模式设计一个程序，模拟这个自助点餐系统的功能。

输入描述
第一行是一个整数 n（1 ≤ n ≤ 100），表示点单的数量。
接下来的 n 行，每行包含一个字符串，表示点餐的饮品名称。

输出描述
输出执行完所有点单后的制作情况，每行输出一种饮品的制作情况。如果制作完成，输出 "XXX is ready!"，其中 XXX 表示饮品名称。

输入示例
4
MilkTea
Coffee
Cola
MilkTea

输出示例
MilkTea is ready!
Coffee is ready!
Cola is ready!
MilkTea is ready!
*/

type Command interface {
	execute()
}

type OrderCommand struct {
	drink      string
	drinkMaker DrinkMaker
}

func (o OrderCommand) execute() {
	o.drinkMaker.action(o.drink)
}

// OrderMachine 调用者
// 可以同时维护一个undo队列，Command定义Undo()接口，实现撤销操作
type OrderMachine struct {
	commands []Command
}

func (o OrderMachine) order() {
	for _, command := range o.commands {
		command.execute()
	}
}

// DrinkMaker 接收者
type DrinkMaker struct {
}

func (d DrinkMaker) action(drink string) {
	fmt.Printf("%s is ready\n", drink)
}

func main() {
	orderMachine := OrderMachine{}
	var count int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var drink string
		fmt.Scan(&drink)
		command := OrderCommand{
			drink:      drink,
			drinkMaker: DrinkMaker{},
		}
		orderMachine.commands = append(orderMachine.commands, command)
		orderMachine.order()
	}
}
