package main

import (
	"fmt"
	"time"
)

/**
2.下面代码是否能编译通过？如果通过，输出什么？
*/
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("not empty interface")
}

func one() {
	var x *int = nil
	Foo(x)

	/**
	参考答案及解析：non-empty interface
	考点：interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
	这里的 x 的动态类型是 *int，所以 x 不为 nil。
	*/
}

/**
3.下面代码输出什么？
*/
func two() {
	ch := make(chan int, 100)
	// A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	// B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a = ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(10)

	/**
	最终会被引起 panic，
	两个 goroutine 启动后还没来得及执行，主协程就把 ch close 了，A 再往 ch 写数据时会触发 panic
	*/
}

func main() {
	one()
	two()
}
