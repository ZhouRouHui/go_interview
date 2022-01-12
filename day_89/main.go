package main

import (
	"fmt"
)

/**
1.下面这段代码能通过编译吗？请简要说明。
*/
func one() {
	v := []int{1, 2, 3}
	for i, n := 0, len(v); i < n; i++ {
		v = append(v, i)
	}
	fmt.Println(v)

	/**
	参考答案及解析：能编译通过，输出 [1 2 3 0 1 2]。for 循环开始的时候，终止条件只会计算一次。
	*/
}

/**
2.下面代码输出什么？
*/
type P *int
type Q *int

func two() {
	var p P = new(int)
	*p += 8
	fmt.Printf("%p\n", p) // 0xc0000ae050
	var x *int = p
	fmt.Printf("%p\n", x) // 0xc0000ae050
	var q Q = x
	*q++
	fmt.Printf("%p\n", q) // 0xc0000ae050
	fmt.Println(*p, *q)

	/**
	参考答案及解析：9 9。指针变量指向相同的地址。
	*/
}

func main() {
	one()
	two()
}
