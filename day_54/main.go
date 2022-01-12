package main

import "fmt"

/**
1.下面的代码有什么问题？
*/
type N int

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v: %p, %v\n", n, *n)
}

func one() {
	var a N = 25
	fmt.Printf("%p, %v\n", &a, a) // 0xc00001e0a0, 25
	p := &a
	fmt.Printf("%p, %v\n", &p, p) // 0xc00000e030, 0xc00001e0a0
	p1 := &p
	fmt.Printf("%p, %v\n", &p1, p1) // 0xc00000e038, 0xc00000e030

	p.value()
	fmt.Printf("%p, %v\n", &a, a) // 0xc0000ae008, 25
	p.pointer()
	fmt.Printf("%p, %v\n", &a, a) // 0xc0000ae008, 26

	//p1.value()
	//p1.pointer()

	/**
	参考答案及解析：编译错误：

	calling method value with receiver p1 (type **N) requires explicit dereference
	calling method pointer with receiver p1 (type **N) requires explicit dereference

	不能使用多级指针调用方法。
	*/
}

/**
2.下面的代码输出什么？
*/
type M int

func (m M) test() {
	fmt.Println(m)
}

func two() {
	var n M = 10
	fmt.Println(n)

	n++
	f1 := M.test
	f1(n)

	n++
	f2 := (*M).test
	f2(&n)

	/**
	参考答案及解析：10 11 12。知识点：方法表达式。通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。
	类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以。

	还可以直接使用方法表达式调用：

	func main()  {
	    var n N = 10

	    fmt.Println(n)

	    n++
	    N.test(n)

	    n++
	    (*N).test(&n)
	}
	引自：《Go语言学习笔记》· 方法
	*/
}

func main() {
	one()
	two()
}
