package main

import "fmt"

func main() {
	//写入defer关键字 java的finnally 压栈输出
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
