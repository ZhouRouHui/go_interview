package main

import "fmt"

/**
2.下面的代码输出什么？
*/
func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func one() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())
	i := f()
	fmt.Println(i)

	/**
	参考答案及解析：7 6 8。
	知识点：匿名函数、defer()。defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出。

	引自《Go语言101》
	*/
}

func main() {
	one()
}
