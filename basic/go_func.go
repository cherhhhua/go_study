package main

import "fmt"

/*
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
*/
func swap(x, y string) (string, string) {
	return y, x
}

/* 定义相互交换值的函数  形参不改变原来的值*/
func swap1(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/* 定义交换值函数  引用传递，地址变化了原来的值也变了*/
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func main() {
	a, b := swap("100", "200")
	c := 10

	fmt.Println(a, b)
	fmt.Printf("变量的地址：%x\n", &c)
}
