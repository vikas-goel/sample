package main

import (
	"fmt"
)

type geometry interface {
	area() int
	perim() int

}

func measure(g geometry) {
		fmt.Println("Area:", g.area(), "Perimeter:", g.perim())
}

type rectangle struct {
	width, height int
	gen interface{}
}

func (r *rectangle) area() int {
	return r.width*r.height
}

func (r *rectangle) perim() int {
	return 2*(r.width+r.height)
}

func (r *rectangle) setWidth(w int) {
	r.width = w
}

func (r *rectangle) getWidth() int {
	return r.width
}

func (r rectangle) setHeight(h int) {
	r.height = h
}

func area(r *rectangle) int {
	return r.width*r.height
}

func getArea(r rectangle) int {
	return r.width*r.height
}

func main() {
	r := &rectangle{2, 4, 3}
	fmt.Println(r.area())

	r.setWidth(5)
	r.setHeight(10)
	fmt.Println(r.area())
	fmt.Println(area(r))
	fmt.Println(getArea(*r))

	r1 := rectangle{height: 3, width: 5}
	fmt.Println(r1.area())
	fmt.Println(r1.getWidth())

	fmt.Printf("&r:%p, &r1:%p\n", &r, &r1)
	measure(r)
	//measure(&r1)

	switch r.gen.(type) {
	case rectangle:
		fmt.Println("Rectangle")
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown")
	}
}
