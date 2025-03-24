package designpattern

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
题目描述
小明家有两个工厂，一个用于生产圆形积木，一个用于生产方形积木，请你帮他设计一个积木工厂系统，记录积木生产的信息。

输入描述
输入的第一行是一个整数 N（1 ≤ N ≤ 100），表示生产的次数。
接下来的 N 行，每行输入一个字符串和一个整数，字符串表示积木的类型。积木类型分为 "Circle" 和 "Square" 两种。整数表示该积木生产的数量

输出描述
对于每个积木，输出一行字符串表示该积木的信息。

输入示例
3
Circle 1
Square 2
Circle 1

输出示例
Circle Block
Square Block
Square Block
Circle Block

提示信息
在示例中，积木工厂生产了4块积木，其中有2块是圆形积木，2块是方形积木。根据输入的类型，每块积木的信息被输出到控制台。
*/

type Block interface {
	Print()
}

type Circle struct {
}

func (c Circle) Print() {
	fmt.Println("Circle Block")
}

type Square struct {
}

func (s Square) Print() {
	fmt.Println("Square Block")
}

type BlockFactory interface {
	Produce() Block
}

type CircleFactory struct {
}

func (c CircleFactory) Produce() Block {
	return &Circle{}
}

type SquareFactory struct {
}

func (s SquareFactory) Produce() Block {
	return Square{}
}

type FactorySystem struct {
	CircleFactory BlockFactory
	SquareFactory BlockFactory
	Blocks        []Block
}

func NewFactorySystem() FactorySystem {
	return FactorySystem{
		CircleFactory: CircleFactory{},
		SquareFactory: SquareFactory{},
	}
}

func (f *FactorySystem) Produce(blockType string, num int) {
	if blockType == "Circle" {
		for i := 0; i < num; i++ {
			f.Blocks = append(f.Blocks, f.CircleFactory.Produce())
		}
		return
	}
	if blockType == "Square" {
		for i := 0; i < num; i++ {
			f.Blocks = append(f.Blocks, f.SquareFactory.Produce())
		}
	}
}

func (f *FactorySystem) Review() {
	for _, block := range f.Blocks {
		block.Print()
	}
}

func main() {
	factorySystem := NewFactorySystem()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < count && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		num, _ := strconv.Atoi(arr[1])
		factorySystem.Produce(arr[0], num)
	}
	factorySystem.Review()
}

/*
func (f *FactorySystem) Produce(blockType string, num int)
因为方法中会对结构体中的数组做append操作，所以方法接收者必须是指针类型，否则修改无法生效。

接收者：表示函数是关联并作用到某个实例上的，相当于this。此时函数升级为方法。
值接收者：在方法调用时创建该值的副本，因此对副本的修改不会影响原始值。
指针接收者：在方法调用时直接使用原始值。因此对指针的修改会影响原始值。
使用场景：
1. 需要在方法中修改结构体的字段时；
2. 需要在方法中避免创建副本如大对象的开销时；
3. 需要在方法中共享结构体的状态时。

函数：独立的一段逻辑；
方法：跟接收者关联。

Go中没有类和成员变量的概念，取代的是结构体和字段。
*/
