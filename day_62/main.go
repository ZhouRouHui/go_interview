package main

import "fmt"

/**
1.下面这段代码输出什么？
*/
func one() {
	nil := 123
	fmt.Println(nil)
	//var _ map[string]int = nil

	/**
	参考答案及解析：var _ map[string]int = nil。
	当前作用域中，预定义的 nil 被覆盖，此时 nil 是 int 类型值，不能赋值给 map 类型。
	*/
}

/**
2.下面代码输出什么？
*/
func two() {
	var x int8 = -128
	var y = x / -1
	fmt.Printf("%T, %v", y, y)

	/**
	参考答案及解析：-128。因为溢出。
	*/
}

func main() {
	one()
	two()
}
