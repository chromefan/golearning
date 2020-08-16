package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 9
	w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Printf("%v", w)
	w = Wheel{
		Circle: Circle{
			Center:  Point{X: 16, Y: 18},
			Radius: 10,
		},
		Spokes: 40,
	}
	fmt.Printf("%v", w)

}
