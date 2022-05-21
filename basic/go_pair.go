package main

import (
	"fmt"
	"io"
	"os"
)

/*
反射，就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type）
，在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type）
，只有interface类型才有反射一说。
*/

func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
