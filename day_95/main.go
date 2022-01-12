package main

import (
	"fmt"
	"time"
)

/**
1.下面这段代码能通过编译吗？请简要说明。
*/
type foo struct{ Val int }

type bar struct{ Val int }

func one() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	fmt.Println(&a, &b)
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)

	/**
	参考答案及解析：false true true。这道题唯一有疑问的地方就在第一个比较，Go 语言里没有引用变量，
	每个变量都占用一个惟一的内存位置，所以第一个比较输出 false。这个知识点在《Go 语言没有引用传递》有介绍。
	*/
}

/**
2.下面代码输出什么？
*/
func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func two() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)

	/**
	参考答案及解析：1、2随机输出，因为两个方法都可以正常执行，而 select 里除了 default 之外是随机的。
	*/
}

func main() {
	one()
	two()
}
