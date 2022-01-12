package main

import "fmt"

/**
1.下面代码有什么问题？
*/
func one() {
	const x = 123
	const y = 1.23
	fmt.Println(x)

	/**
	参考答案及解析：编译可以通过。
	知识点：常量。常量是一个简单值的标识符，在程序运行时，不会被修改的量。不像变量，常量未使用是能编译通过的。const 可定以全局和局部常量。
	*/
}

/**
2.下面代码输出什么？
*/
const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func two() {
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)

	/**
	参考答案及解析：
	    uint16 120
	    string abc
	知识点：常量。
	常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	*/
}

func main() {
	one()
	two()
}
