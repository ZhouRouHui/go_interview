package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
	m
)

/**
2.下面这段代码能否编译通过？如果可以，输出什么？
*/
func one() {
	fmt.Println(x, y, z, k, p, m) // 0 2 zz zz 5 6
}

/**
2. 给一个变量初始为 nil 的方式
*/
func two() {
	var a interface{} = nil
	var b error = nil
	fmt.Println(a, b)

	/**
	nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，它是一种内置接口类型，看下方贴出的源码就知道，所以 D 是对的。
	type error interface {
	    Error() string
	}
	*/
}

func main() {
	one()
	two()
}
