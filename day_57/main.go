package main

import "fmt"

/**
1.下面哪一行代码会 panic，请说明原因？
*/
func one() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = x == y
	//_ = y == y

	/**
	参考答案及解析：_ = y == y
	切片不可比较
	*/
}

/**
2.下面代码输出什么？
*/
var o = fmt.Println

func two() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- 1:
			o(3)
		}
	}
	/**
	参考答案及解析：321。
	第一次循环，写操作已经准备好，执行 o(3)，输出 3；
	第二次，读操作准备好，执行 o(2)，输出 2 并将 c 赋值为 nil；
	第三次，由于 c 为 nil，走的是 default 分支，输出 1。

	两题均引自：《Go语言101》
	*/
}

func main() {
	one()
	two()
}
