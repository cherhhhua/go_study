package main

import (
	"fmt"
	"unsafe"
)

/*
常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
const identifier [type] = value
常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过。
*/

// const b string = "abc"

// const c = "abc"

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	CategoryBooks    = iota // 0
	CategoryHealth          // 1
	CategoryClothing        // 2
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	// const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println(a, b, c)

	println(Unknown, Female, Male)

	println(CategoryBooks, CategoryHealth, CategoryClothing)

	println(IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)

	println(IgEggs | IgChocolate | IgShellfish)

	println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
}
