package main

import "C"
import (
	"fmt"
	"testing"
)

/**
1.关于 cap 函数适用下面哪些类型？
*/
func one() {
	/*
		A. 数组；
		B. channel;
		C. map；
		D. slice；
	*/

	/**
	参考答案即解析：ABD。cap() 函数的作用：

	array 返回数组的元素个数；
	slice 返回 slice 的最大容量；
	channel 返回 channel 的容量；
	*/
}

/**
2.下面代码输出什么？
*/
func hello(num ...int) {
	num[0] = 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func two() {
	t := &testing.T{}
	Test13(t)

	/**
	参考答案及解析：18。可变函数是指针传递。相当于切片
	*/
}

func main() {
	one()
	two()
}
