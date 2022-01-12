package main

import "fmt"

/**
1.下面代码输出什么？
*/
type N int

func (i *N) test() {
	fmt.Println(*i)
}

func one() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()

	/**
	参考答案及解析：13 13 13。
	知识点：方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针。
	*/
}

/**
2.下面哪一行代码会 panic，请说明原因？
*/
func two() {
	var m map[int]bool
	// 从一个 nil 的 map 取值不会 panic，给一个 nil 的 map 设置值会 panic
	_ = m[123]
	var p *[5]string
	for range p {
		_ = len(p)
	}
	var s []int
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
	/**
	参考答案及解析：s, s[0] = []int{1, 2}, 9。
	因为左侧的 s[0] 中的 s 为 nil。

	引自：《Go语言101》
	*/
}

func main() {
	one()
	two()
}
