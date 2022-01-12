package main

import "fmt"

/**
1.下面这段代码输出什么？为什么？
*/

const (
	a = iota
	b = iota
)
const (
	c = "c"
	d = "d"
	e = "e"
	f = iota
	g = iota
)

func main() {
	fmt.Println(a, b, c, d, e, f, g)
	/**
	参考答案及解析：0 1 c d e 3 4。知识点：iota 的用法。

	iota 是 golang 语言的常量计数器，只能在常量的表达式中使用。

	iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次。如果 iota 不是在 const 的第一行，
	那么 iota 之前的都算在计数器以内
	*/
}
