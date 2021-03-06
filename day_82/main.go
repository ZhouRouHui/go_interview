package main

import "fmt"

/**
1.下面的代码输出什么？
*/
func one() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
		}
		if m == -m {
			count++
		}
	}
	fmt.Println(count)

	/**
	参考答案及解析：4。知识点：数值溢出。当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名。

	引自《Go语言101》
	*/
}

/**
2.下面代码输出什么？
*/
const (
	azero = iota
	aone  = iota
)

const (
	info  = "msg"
	bzero = iota
	bone  = iota
)

func two() {
	fmt.Println(azero, aone)
	fmt.Println(bzero, bone)

	/**
	参考答案及解析：
	0 1
	1 2。
	知识点：iota 的使用。这道题易错点在 bzero、bone 的值，在一个常量声明代码块中，如果 iota 没出现在第一行，则常量的初始值就是非 0 值。
	*/
}

func main() {
	one()
	two()
}
