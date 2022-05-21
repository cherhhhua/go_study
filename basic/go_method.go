package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Add(another Point) Point {
	return Point{p.X + another.X, p.Y + another.Y}
}

func (p Point) Sub(another Point) Point {
	return Point{p.X - another.X, p.Y - another.Y}
}

func (p Point) Print() {
	fmt.Printf("{%f, %f}\n", p.X, p.Y)
}

//定义一个Point切片类型 Path
type Path []Point

//方法的接收器 是Path类型数据, 方法的选择器是TranslateBy(Point, bool)
func (path Path) TranslateBy(another Point, add bool) {
	var op func(p, q Point) Point //定义一个 op变量 类型是方法表达式 能够接收Add,和 Sub方法
	if add == true {
		op = Point.Add //给op变量赋值为Add方法
	} else {
		op = Point.Sub //给op变量赋值为Sub方法
	}

	for i := range path {
		//调用 path[i].Add(another) 或者 path[i].Sub(another)
		path[i] = op(path[i], another)
		path[i].Print()
	}
}
func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFormP := p.Distance   // 方法值(相当于C语言的函数地址,函数指针)
	fmt.Println(distanceFormP(q)) // "5"
	fmt.Println(p.Distance(q))    // "5"

	//实际上distanceFormP 就绑定了 p接收器的方法Distance

	distanceFormQ := q.Distance   //
	fmt.Println(distanceFormQ(p)) // "5"
	fmt.Println(q.Distance(p))    // "5"

	//实际上distanceFormQ 就绑定了 q接收器的方法Distance
	fmt.Println("------------------")

	distance1 := Point.Distance //方法表达式, 是一个函数值(相当于C语言的函数指针)
	fmt.Println(distance1(p, q))
	fmt.Printf("%T\n", distance1) //%T表示打出数据类型 ,这个必须放在Printf使用

	distance2 := (*Point).Distance //方法表达式,必须传递指针类型
	distance2(&p, q)
	fmt.Printf("%T\n", distance2)
	fmt.Println("------------------")

	points := Path{
		{10, 10},
		{11, 11},
	}

	anotherPoint := Point{5, 5}

	points.TranslateBy(anotherPoint, false)

	fmt.Println("------------------")

	points.TranslateBy(anotherPoint, true)
}

/*
在一个包的API需要一个函数值、且调用方希望操作的是某一个绑定了对象的方法的话，方法"值"会非常实用.


这个Distance实际上是指定了Point对象为接收器的一个方法func (p Point) Distance()，
但通过Point.Distance得到的函数需要比实际的Distance方法多一个参数，
即其需要用第一个额外参数指定接收器，后面排列Distance方法的参数。
看起来本书中函数和方法的区别是指有没有接收器，而不像其他语言那样是指有没有返回值。
*/
