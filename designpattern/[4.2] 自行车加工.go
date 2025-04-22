package designpattern

import "fmt"

/**
题目描述
小明家新开了一家自行车工厂，用于使用自行车配件（车架 frame 和车轮 tires ）进行组装定制不同的自行车，包括山地车和公路车。
山地车使用的是Aluminum Frame（铝制车架）和 Knobby Tires（可抓地轮胎），公路车使用的是 Carbon Frame （碳车架）和 Slim Tries。
现在它收到了一笔订单，要求定制一批自行车，请你使用【建造者模式】告诉小明这笔订单需要使用那些自行车配置吧。

输入描述
输入的第一行是一个整数 N（1 ≤ N ≤ 100），表示订单的数量。

接下来的 N 行，每行输入一个字符串，字符串表示客户的自行车需求。

字符串可以包含关键词 "mountain" 或 "road"，表示客户需要山地自行车或公路自行车。

输出描述
对于每笔订单，输出该订单定制的自行车配置。

输入示例
3
mountain
road
mountain

输出示例
Aluminum Frame Knobby Tires
Carbon Frame Slim Tires
Aluminum Frame Knobby Tires

提示信息
在本例中：产品为自行车，可以有两个建造者：山地车建造者和公路车建造者。
*/

type Bike struct {
	Frame string
	Tires string
}

func (b Bike) String() string {
	return fmt.Sprintf("%s %s\n", b.Frame, b.Tires)
}

type Builder interface {
	BuildFrame()
	BuildTires()
	GetBike() *Bike
}

type MountainBikeBuilder struct {
	MountainBike *Bike
}

func NewMountainBikeBuilder() *MountainBikeBuilder {
	return &MountainBikeBuilder{
		MountainBike: &Bike{},
	}
}

func (m *MountainBikeBuilder) BuildFrame() {
	m.MountainBike.Frame = "Aluminum Frame"
}

func (m *MountainBikeBuilder) BuildTires() {
	m.MountainBike.Tires = "Knobby Tires"
}

func (m *MountainBikeBuilder) GetBike() *Bike {
	return m.MountainBike
}

type RoadBikeBuilder struct {
	RoadBike *Bike
}

func NewRoadBikeBuilder() *RoadBikeBuilder {
	return &RoadBikeBuilder{
		RoadBike: &Bike{},
	}
}

func (r *RoadBikeBuilder) BuildFrame() {
	r.RoadBike.Frame = "Carbon Frame"
}

func (r *RoadBikeBuilder) BuildTires() {
	r.RoadBike.Tires = "Slim Tries"
}

func (r *RoadBikeBuilder) GetBike() *Bike {
	return r.RoadBike
}

type BikeDirector struct{}

func (b *BikeDirector) construct(builder Builder) *Bike {
	builder.BuildFrame()
	builder.BuildTires()
	return builder.GetBike()
}

func main() {
	var count int
	_, _ = fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var bikeType string
		_, _ = fmt.Scan(&bikeType)
		var builder Builder
		switch bikeType {
		case "mountain":
			builder = NewMountainBikeBuilder()
		case "road":
			builder = NewRoadBikeBuilder()
		}
		director := &BikeDirector{}
		bike := director.construct(builder)
		fmt.Println(bike)
	}
}

/**
结构体实现一个方法：
func (x XXX) String() string {
	return ""
}
可以方便地打印出return的值，达到序列化效果：
fmt.Println(x)
*/
