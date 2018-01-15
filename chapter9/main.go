package main

import ("fmt"; "math")

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r*r
}

func newcircleArea(c Circle) float64 {
	return math.Pi * c.r*c.r
}

func pcircleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}


func (c *Circle) mcircleArea() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) mrectangleArea() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	var c Circle
	c = Circle{0, 0, 10}
	//fmt.Println(c.x, c.y, c.r)
	fmt.Println(newcircleArea(c))

	d := new(Circle) //default 0
	//fmt.Println(d.x,d.y,d.r)
	d.x=0
	d.y=0
	d.r=20
	fmt.Println(pcircleArea(d))
	
	e := Circle{x: 0, y: 0, r: 20}
	fmt.Println(pcircleArea(&e))

	f := Circle{0,0,20}
	fmt.Println(f.mcircleArea())

	g := Rectangle{0,0,10,10}
	fmt.Println(g.mrectangleArea())

	h := Person{"New"}
	h.Talk()


	a := new(Android)
	a.Person.Talk()
	a.Talk()

}