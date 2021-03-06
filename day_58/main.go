package main

import "fmt"

/**
1.下面代码输出什么？
*/
type T struct {
	x int
	y *int
}

func one() {
	i := 20
	t := T{10, &i}

	p := &t.x
	*p++
	*p--

	t.y = p
	fmt.Println(*t.y)

	/**
	参考答案及解析：10。
	知识点：运算符优先级。
	如下规则：
	递增运算符 ++ 和递减运算符 – 的优先级低于解引用运算符 * 和取址运算符 &，
	解引用运算符和取址运算符的优先级低于选择器 . 中的属性选择操作符。


	*/
}

/**
2.下面哪一行代码会 panic，请说明原因？
*/
func two() {
	x := make([]int, 2, 10)
	_ = x[6:10]
	//_ = x[6:]
	_ = x[2:]

	/**
	参考答案：_ = x[6:]，
	截取符号 [i:j]，如果 j 省略，默认是原切片或者数组的长度，x 的长度是 2，小于起始下标 6 ，所以 panic。

	两题均引自：《Go语言101》
	*/
}

func main() {
	one()
	two()
}
