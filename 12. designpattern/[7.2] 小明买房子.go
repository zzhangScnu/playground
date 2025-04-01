package designpattern

import "fmt"

/**
题目描述
小明想要购买一套房子，他决定寻求一家房屋中介来帮助他找到一个面积超过100平方米的房子，只有符合条件的房子才会被传递给小明查看。

输入描述
第一行是一个整数 N（1 ≤ N ≤ 100），表示可供查看的房子的数量。
接下来的 N 行，每行包含一个整数，表示对应房子的房屋面积。

输出描述
对于每个房子，输出一行，表示是否符合购房条件。如果房屋面积超过100平方米，输出 "YES"；否则输出 "NO"。

输入示例
3
120
80
110

输出示例
YES
NO
YES
*/

type House interface {
	View()
}

type Person struct {
}

func (p *Person) View() {
	fmt.Println("YES")
}

type Proxy struct {
	Person Person
}

func (p *Proxy) View(area int) {
	if area > 100 {
		p.Person.View()
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var count int
	fmt.Scan(&count)
	proxy := &Proxy{Person: Person{}}
	var area int
	for i := 0; i < count; i++ {
		fmt.Scan(&area)
		proxy.View(area)
	}
}
