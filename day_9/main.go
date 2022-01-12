package main

import "fmt"

func hello(num ...int) {
	fmt.Printf("%p\n", num)
	num[0] = 18
}

func test(num int) {
	fmt.Printf("%p\n", &num)
	num = 22
}

func main() {
	i := []int{5, 6, 7}
	fmt.Printf("%p\n", i)
	hello(i...)
	fmt.Println(i[0]) // 18, 切片是引用类型，直接改元数据

	a := 2
	fmt.Printf("%p\n", &a)
	test(a)
	fmt.Println(a) // 2，标准类型是值类型，改的是变量的副本
}
