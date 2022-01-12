package main

import "fmt"

/**
2.下面这段代码输出什么？
*/
func one() {
	i := 65
	fmt.Println(string(i))

	/**
	参考答案及解析：A。
	知识点：UTF-8 编码中，十进制数字 65 对应的符号是 A。
	*/
}

/**
3.下面这段代码输出什么？
*/
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func two() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())

	/**
	参考答案及解析：13 23。
	知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
	*/
}

func main() {
	one()
	two()
}
