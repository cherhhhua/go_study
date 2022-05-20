package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
}

// 运行结果:
// type:  float64
// value:  1.2345

/*
●反射可以大大提高程序的灵活性，使得interface{}有更大的发挥余地
	○反射必须结合interface才玩得转
	○变量的type要是concrete type的（也就是interface变量）才有反射一说
●反射可以将“接口类型变量”转换为“反射类型对象”
	○反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
●反射可以将“反射类型对象”转换为“接口类型变量
	○reflect.value.Interface().(已知的类型)
	○遍历reflect.Type的Field获取其Field
●反射可以修改反射类型对象，但是其值必须是“addressable”
	○想要利用反射修改对象状态，前提是 interface.data 是 settable,即 pointer-interface
●通过反射可以“动态”调用方法
●因为Golang本身不支持模板，因此在以往需要使用模板的场景下往往就需要使用反射(reflect)来实现
*/
