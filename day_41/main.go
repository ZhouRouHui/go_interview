package main

import (
	"fmt"
)

/**
下面这段代码输出什么
*/
func one() {
	s := []int{2: 2, 3, 0: 1}
	fmt.Println(s)

	/**
	参考答案及解析：输出[1 0 2 3]，字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一，所以输出[1 0 2 3]，而不是[1 3 2]。
	*/
}

/**
3.下面这段代码输出什么？
*/
func incr(p *int) int {
	*p++
	return *p
}

func two() {
	v := 1
	incr(&v)
	fmt.Println(v)

	/**
	参考答案及解析：2。知识点：指针。p 是指针变量，指向变量 v，*p++操作的意思是取出变量 v 的值并执行加一操作，所以 v 的最终值是 2。
	*/
}

func main() {
	one()
	two()
}
