package designpattern

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/**
题目描述
小明去了一家大型商场，拿到了一个购物车，并开始购物。
请你设计一个购物车管理器，记录商品添加到购物车的信息（商品名称和购买数量），并在购买结束后打印出商品清单。
（在整个购物过程中，小明只有一个购物车实例存在）。

输入描述
输入包含若干行，每行包含两部分信息，分别是商品名称和购买数量。商品名称和购买数量之间用空格隔开。

输出描述
输出包含小明购物车中的所有商品及其购买数量。每行输出一种商品的信息，格式为 "商品名称 购买数量"。

输入示例
Apple 3
Banana 2
Orange 5

输出示例
Apple 3
Banana 2
Orange 5

提示信息
本道题目请使用单例设计模式：
使用私有静态变量来保存购物车实例。
使用私有构造函数防止外部直接实例化。
*/

var cart *Cart

var once sync.Once

type Cart struct {
	lock  sync.Mutex
	names []string
	items map[string]int
}

func newCart() *Cart {
	once.Do(func() {
		cart = &Cart{
			items: make(map[string]int),
		}
	})
	return cart
}

func (c *Cart) add(name string, num int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if _, ok := c.items[name]; !ok {
		c.names = append(c.names, name)
	}
	c.items[name] += num
}

func (c *Cart) output() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, name := range c.names {
		fmt.Printf("%s %d\n", name, c.items[name])
	}
}

func main() {
	cart = newCart()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content := scanner.Text()
		if content == "" {
			break
		}
		arr := strings.Split(content, " ")
		if len(arr) < 2 {
			continue
		}
		name := arr[0]
		num, err := strconv.Atoi(arr[1])
		if err != nil {
			continue
		}
		cart.add(name, num)
	}
	cart.output()
}

/**
1. sync.Once包实现双重检查加锁的懒加载；
2. sync.Mutex避免临界区变量的竞态访问，包括读和写；
3. sync包提供同步机制，变量禁止拷贝。如sync.Mutex如果被拷贝，原始锁和拷贝锁将拥有各自独立的状态，无法保护共享资源；
4. 用slice & map的组合，实现固定顺序 + 去重，即Java中的LinkedListMap；
5. bufio.Scanner可以实现动态读取输入，即ACM模式的标准输入。
*/
