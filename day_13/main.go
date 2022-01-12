package main

import "fmt"

/**
2.下面这段代码输出什么？
*/
func hello(i int) {
	fmt.Println(i)
}

func one() {
	i := 5
	defer hello(i)
	i = i + 10

	/**
	参考答案及解析：5。这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
	*/
}

/**
3.下面这段代码输出什么？
*/
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func two() {
	t := Teacher{}
	t.ShowA()

	/**
	参考答案及解析：
	showA
	showB。

	知识点：结构体嵌套。这道题可以结合第 12 天的第三题一起看，Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法.
	*/
}

func main() {
	one()
	two()
}
