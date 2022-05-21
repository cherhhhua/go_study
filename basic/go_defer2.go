package main

import "fmt"

//return先于defer执行
func deferfunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnfunc() int {
	fmt.Println("return func called...")
	return 0
}
func returnAndDefer() int {
	defer deferfunc()
	return returnfunc()
}

func main() {
	returnAndDefer()
}
