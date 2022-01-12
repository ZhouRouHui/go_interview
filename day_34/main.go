package main

import "fmt"

type Integer int

func (i *Integer) Add(b Integer) Integer {
	return *i + b
}

// 下面这种方式也可以
//func (i Integer) Add(b Integer) Integer {
//	return i + b
//}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b) // i.(*Integer) 类型断言，认定并转换 i 为 *Integer 类型
	fmt.Println(sum)
}
