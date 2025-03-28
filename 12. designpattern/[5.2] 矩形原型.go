package designpattern

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type Rectangle struct {
	Color  string
	Width  int
	Height int
}

func (r *Rectangle) Clone() Prototype {
	return &Rectangle{
		Color:  r.Color,
		Width:  r.Width,
		Height: r.Height,
	}
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Color: %s, Width: %d, Height: %d", r.Color, r.Width, r.Height)
}

func main() {
	var size int
	fmt.Scan(&size)
	for i := 0; i < size; i++ {
		var color string
		var width, height int
		fmt.Scan(&color, &width, &height)
		prototype := &Rectangle{
			Color:  color,
			Width:  width,
			Height: height,
		}
		rectangle := prototype.Clone().(*Rectangle)
		fmt.Println(rectangle.String())
	}
}
