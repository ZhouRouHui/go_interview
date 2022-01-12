package main

import "fmt"

/**
2.下面这段代码输出什么？
*/
func incr(p *int) int {
	*p++
	return *p
}

func one() {
	p := 1
	incr(&p)
	fmt.Println(p)

	/**
	参考答案及解析：2。
	知识点：指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 one() 函数的变量 p 的地址。*p++ 是将该地址的值执行一个自增操作，
	incr() 返回自增后的结果。
	*/
}

func main() {
	one()
}
