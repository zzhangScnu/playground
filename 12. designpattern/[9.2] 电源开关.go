package designpattern

import "fmt"

/**
题目描述
小明家的电源总开关控制了家里的三个设备：空调、台灯和电视机。
每个设备都有独立的开关密码，分别用数字1、2和3表示。
即输入1时，空调关闭，输入2时，台灯关闭，输入3时，电视机关闭，当输入为4时，表示要关闭所有设备。
请你使用外观模式编写程序来描述电源总开关的操作。

输入描述
第一行是一个整数 N（1 <= N <= 100），表示后面有 N 行输入。
接下来的 N 行，每行包含一个数字，表示对应设备的开关操作（1表示关闭空调，2表示关闭台灯，3表示关闭电视机，4表示关闭所有设备）。

输出描述
输出关闭所有设备后的状态，当输入的数字不在1-4范围内时，输出Invalid device code.

输入示例
4
1
2
3
4

输出示例
Air Conditioner is turned off.
Desk Lamp is turned off.
Television is turned off.
All devices are off.
*/

type AirConditioner struct {
}

func (c AirConditioner) Turnoff() {
	fmt.Println("Air Conditioner is turned off.")
}

type Lamp struct {
}

func (l Lamp) Turnoff() {
	fmt.Println("Desk Lamp is turned off.")
}

type Television struct {
}

func (t Television) Turnoff() {
	fmt.Println("Television is turned off.")
}

type Switch struct {
	airConditioner AirConditioner
	lamp           Lamp
	television     Television
}

func NewSwitch() Switch {
	return Switch{
		airConditioner: AirConditioner{},
		lamp:           Lamp{},
		television:     Television{},
	}
}

func (s Switch) Turnoff() {
	fmt.Println("All devices are off.")
}

func (s Switch) Control(password int) {
	switch password {
	case 1:
		s.airConditioner.Turnoff()
	case 2:
		s.lamp.Turnoff()
	case 3:
		s.television.Turnoff()
	case 4:
		s.Turnoff()
	default:
		fmt.Println("Invalid device code.")
	}
}

func main() {
	var count int
	fmt.Scan(&count)
	s := NewSwitch()
	for i := 0; i < count; i++ {
		var password int
		fmt.Scan(&password)
		s.Control(password)
	}
}
