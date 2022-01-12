package main

import (
	"fmt"
	"sync"
)

/**
1.下面代码输出什么？
*/
var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}
func one() {
	go f()
	c <- 0
	println(a)

	/**
	参考答案及解析：1。能正确输出，不过主协程会阻塞 f() 函数的执行。
	*/
}

/**
2.下面代码输出什么？请简要说明。
A. 不能编译；
B. 输出 1, 1；
C. 输出 1, 2；
D. fatal error；
*/
type MyMutex struct {
	count int
	sync.Mutex
}

func two() {
	var mu MyMutex
	mu.Lock()
	fmt.Printf("%p\n", &mu.Mutex)
	var mu1 = mu
	fmt.Printf("%p\n", &mu1.Mutex)
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)

	/**
	参考答案及解析：D。加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁。
	*/
}

func main() {
	one()
	two()
}
