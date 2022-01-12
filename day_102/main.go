package main

import (
	"fmt"
	"sync"
)

/**
2.下面的代码输出什么？请简要说明
*/
var mu sync.Mutex
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> A"
	B()
}

func B() {
	chain = chain + " --> B"
	C()
}

func C() {
	// 互斥锁的锁不能重复加锁
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}

func two() {
	chain = "main"
	A()
	fmt.Println(chain)

	/**
	参考答案即解析：D。使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁。
	引自博客《鸟窝》 https://colobu.com/
	*/
}

func main() {
	two()
}
