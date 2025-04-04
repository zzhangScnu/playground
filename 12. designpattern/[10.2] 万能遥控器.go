package designpattern

import "fmt"

/**
题目描述
小明家有一个万能遥控器，能够支持多个品牌的电视。
每个电视可以执行开机、关机和切换频道的操作，请你使用桥接模式模拟这个操作。

输入描述
第一行是一个整数 N（1 <= N <= 100），表示后面有 N 行输入。
接下来的 N 行，每行包含两个数字。第一个数字表示创建某个品牌的遥控和电视，第二个数字表示执行的操作。
其中，0 表示创建 Sony 品牌的遥控和电视，1 表示创建 TCL 品牌的遥控和电视；
2 表示开启电视、3表示关闭电视，4表示切换频道。

输出描述
对于每个操作，输出相应的执行结果。

输入示例
6
0 2
1 2
0 4
0 3
1 4
1 3

输出示例
Sony TV is ON
TCL TV is ON
Switching Sony TV channel
Sony TV is OFF
Switching TCL TV channel
TCL TV is OFF
*/

type TV interface {
	ON()
	OFF()
	SwitchChannel()
}

type SONY struct {
}

func (s SONY) ON() {
	fmt.Println("Sony TV is ON")
}

func (r SONY) OFF() {
	fmt.Println("Sony TV is OFF")
}

func (r SONY) SwitchChannel() {
	fmt.Println("Switching Sony TV channel")
}

type TCL struct {
}

func (t TCL) ON() {
	fmt.Println("TCL TV is ON")
}

func (t TCL) OFF() {
	fmt.Println("TCL TV is OFF")
}

func (t TCL) SwitchChannel() {
	fmt.Println("Switching TCL TV channel")
}

type Remote interface {
	Operate()
}

type OnOperation struct {
	tv TV
}

type OffOperation struct {
	tv TV
}

type SwitchChannelOperation struct {
	tv TV
}

func (oo OnOperation) Operate() {
	oo.tv.ON()
}

func (oo OffOperation) Operate() {
	oo.tv.OFF()
}

func (wco SwitchChannelOperation) Operate() {
	wco.tv.SwitchChannel()
}

func main() {
	var count int
	fmt.Scan(&count)
	var tvType, operationType int
	for i := 0; i < count; i++ {
		var tv TV
		var remote Remote
		fmt.Scan(&tvType, &operationType)
		switch tvType {
		case 0:
			tv = &SONY{}
		case 1:
			tv = &TCL{}
		}
		switch operationType {
		case 2:
			remote = &OnOperation{tv: tv}
		case 3:
			remote = &OffOperation{tv: tv}
		case 4:
			remote = &SwitchChannelOperation{tv: tv}
		}
		remote.Operate()
	}
}
