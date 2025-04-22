package designpattern

import "fmt"

/**
题目描述
小明所在的学校有一个时钟（主题），每到整点时，它就会通知所有的学生（观察者）当前的时间，请你使用观察者模式实现这个时钟通知系统。
注意点：时间从 0 开始，并每隔一个小时更新一次。

输入描述
输入的第一行是一个整数 N（1 ≤ N ≤ 20），表示学生的数量。
接下来的 N 行，每行包含一个字符串，表示学生的姓名。
最后一行是一个整数，表示时钟更新的次数。

输出描述
对于每一次时钟更新，输出每个学生的姓名和当前的时间。

输入示例
2
Alice
Bob
3

输出示例
Alice 1
Bob 1
Alice 2
Bob 2
Alice 3
Bob 3

提示信息
初始时钟时间为0（12:00 AM）。
第一次更新后，时钟变为1（1:00 AM），然后通知每个学生，输出学生名称和时钟点数。
第二次更新后，时钟变为2（2:00 AM），然后再次通知每个学生，输出学生名称和时钟点数
第三次更新后，时钟变为3（3:00 AM），然后再次通知每个学生，输出学生名称和时钟点数。
*/

type Subject interface {
	Register(student Observer)
	Remove(student Observer)
	NotifyAll()
}

type Clock struct {
	Hour      int
	Observers []Observer
}

func (c *Clock) Register(student Observer) {
	c.Observers = append(c.Observers, student)
}

func (c *Clock) Remove(observer Observer) {
	for i, o := range c.Observers {
		if o == observer {
			c.Observers = append(c.Observers[:i], c.Observers[i+1:]...)
			break
		}
	}
}

func (c *Clock) NotifyAll() {
	for time := 0; time < c.Hour; time++ {
		for _, observer := range c.Observers {
			observer.Update((time + 1) % 24)
		}
	}
}

type Observer interface {
	Update(hour int)
}

type Student struct {
	Name string
}

func (c Student) Update(hour int) {
	fmt.Printf("%s %d\n", c.Name, hour)
}

func main() {
	var studentCount, clockTimes int
	fmt.Scan(&studentCount)
	clock := &Clock{}
	for i := 0; i < studentCount; i++ {
		var studentName string
		fmt.Scan(&studentName)
		clock.Register(&Student{Name: studentName})
	}
	fmt.Scan(clockTimes)
	clock.Hour = clockTimes
	clock.NotifyAll()
}
