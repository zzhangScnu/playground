package designpattern

import "fmt"

/**
题目描述
在一个图形编辑器中，用户可以绘制不同类型的图形，包括圆形（CIRCLE）、矩形（RECTANGLE）、三角形（TRIANGLE）等。
现在，请你实现一个图形绘制程序，要求能够共享相同类型的图形对象，以减少内存占用。

输入描述
输入包含多行，每行表示一个绘制命令。每个命令包括两部分：
图形类型（Circle、Rectangle 或 Triangle）
绘制的坐标位置（两个整数，分别表示 x 和 y）

输出描述
对于每个绘制命令，输出相应图形被绘制的位置信息。如果图形是首次绘制，输出 "drawn at"，否则输出 "shared at"。

输入示例
CIRCLE 10 20
RECTANGLE 30 40
CIRCLE 15 25
TRIANGLE 5 15
CIRCLE 10 20
RECTANGLE 30 40

输出示例
CIRCLE drawn at (10, 20)
RECTANGLE drawn at (30, 40)
CIRCLE shared at (15, 25)
TRIANGLE drawn at (5, 15)
CIRCLE shared at (10, 20)
RECTANGLE shared at (30, 40)
*/

type Position struct {
	X int
	Y int
}

type ShapeType string

const (
	CIRCLE    ShapeType = "CIRCLE"
	RECTANGLE ShapeType = "RECTANGLE"
	TRIANGLE  ShapeType = "TRIANGLE"
)

type Shape interface {
	Draw(position Position)
}

type ConcreteShape struct {
	shapeType ShapeType
	created   bool
}

func (c *ConcreteShape) GetCreated() string {
	if c.created {
		return "shared"
	}
	return "drawn"
}

func (c *ConcreteShape) Draw(position Position) {
	fmt.Printf("%s %s at (%d, %d)\n", string(c.shapeType), c.GetCreated(), position.X, position.Y)
}

func NewConcreteShape(shapeType ShapeType) *ConcreteShape {
	return &ConcreteShape{
		shapeType: shapeType,
		created:   false,
	}
}

type ShapeFactory struct {
	shapes map[ShapeType]Shape
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		shapes: make(map[ShapeType]Shape),
	}
}

func (s *ShapeFactory) GetShape(shapeType ShapeType) Shape {
	_, ok := s.shapes[shapeType]
	if !ok {
		s.shapes[shapeType] = NewConcreteShape(shapeType)
	}
	return s.shapes[shapeType]
}

func main() {
	var shapeType string
	var x, y int
	shapeFactory := NewShapeFactory()
	for {
		fmt.Scan(&shapeType, &x, &y)
		if shapeType == "" || x == 0 || y == 0 {
			return
		}
		shape := shapeFactory.GetShape(ShapeType(shapeType))
		shape.Draw(Position{x, y})
		shape.(*ConcreteShape).created = false
	}
}
