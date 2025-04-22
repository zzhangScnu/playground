package designpattern

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
题目描述
小明家新开了两个工厂用来生产家具，一个生产现代风格的沙发和椅子，一个生产古典风格的沙发和椅子，
现在工厂收到了一笔订单，请你帮他设计一个系统，描述订单需要生产家具的信息。

输入描述
输入的第一行是一个整数 N（1 ≤ N ≤ 100），表示订单的数量。
接下来的 N 行，每行输入一个字符串，字符串表示家具的类型。家具类型分为 "modern" 和 "classical" 两种。

输出描述
对于每笔订单，输出字符串表示该订单需要生产家具的信息。
modern订单会输出下面两行字符串
modern chair
modern sofa

classical订单会输出下面两行字符串
classical chair
classical soft

输入示例
3
modern
classical
modern

输出示例
modern chair
modern sofa
classical chair
classical sofa
modern chair
modern sofa

提示信息
在示例中，工厂收到了3笔订单，其中有2笔要求生产modern风格，1笔要求生产classical风格。根据输入的类型，每次订单生产的家具信息被输出到控制台上。
*/

type Chair interface {
	View()
}

type ModernChair struct {
}

func (m ModernChair) View() {
	fmt.Println("modern chair")
}

type ClassicalChair struct {
}

func (c ClassicalChair) View() {
	fmt.Println("classical chair")
}

type Sofa interface {
	View()
}

type ModernSofa struct {
}

func (m ModernSofa) View() {
	fmt.Println("modern sofa")
}

type ClassicalSofa struct {
}

func (c ClassicalSofa) View() {
	fmt.Println("classical sofa")
}

type Factory interface {
	ProduceChair() Chair
	ProduceSofa() Sofa
}

type ModernFactory struct {
}

func (m *ModernFactory) ProduceChair() Chair {
	return ModernChair{}
}

func (m *ModernFactory) ProduceSofa() Sofa {
	return ModernSofa{}
}

type ClassicalFactory struct {
}

func (c *ClassicalFactory) ProduceChair() Chair {
	return ClassicalChair{}
}

func (c *ClassicalFactory) ProduceSofa() Sofa {
	return ClassicalSofa{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < count; i++ {
		scanner.Scan()
		factoryType := scanner.Text()
		var factory Factory
		switch factoryType {
		case "modern":
			factory = &ModernFactory{}
		case "classical":
			factory = &ClassicalFactory{}
		}
		chair := factory.ProduceChair()
		sofa := factory.ProduceSofa()
		chair.View()
		sofa.View()
	}
}
