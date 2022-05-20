package main

import "fmt"

/*
slice
声明一个未指定大小的数组来定义切片
var identifier []type 切片不需要说明长度。
或使用make()函数来创建切片:
var slice1 []type = make([]type, len)
也就是
slice1 := make([]type, len)

make([]T, length, capacity)

s := arr[startIndex:endIndex]
包含左边，不包含右边

append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
*/

func main() {
	var numbers []int
	s := []int{1, 2, 3}

	fmt.Println(s)
	printSlice(s)
	fmt.Println(s[1:])
	printSlice(s[1:])
	fmt.Println(s[:1])
	printSlice(s[:1])

	printSlice(numbers)
	numbers = append(numbers, 1)
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers)
	printSlice(numbers1)
	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
