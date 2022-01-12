package main

import (
	"fmt"
	"sync"
	"time"
)

/**
1.下面代码输出什么？
*/
func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}
func one() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))

	/**
	参考答案及解析：输出 0 40 50。
	知识点：defer 语句与返回值。函数的 return value 不是原子操作，而是在编译器中分解为两部分：返回值赋值 和 return。
	*/
}

/**
2.下面的代码有什么问题？
*/
var mu sync.RWMutex
var count int

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
	fmt.Println(3)
}
func B() {
	time.Sleep(5 * time.Second)
	C()
}
func C() {
	// 读写锁的读锁可以重复加锁，但是当写锁阻塞时，新的读锁无法申请
	fmt.Println(1)
	mu.RLock()
	defer mu.RUnlock()
}

func two() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	fmt.Println(2)
	defer mu.Unlock()
	count++
	fmt.Println(count)

	/**
	参考答案及解析：fatal error。当写锁阻塞时，新的读锁是无法申请的（有效防止写锁饥饿），导致死锁。
	*/
}

func main() {
	one()
	two()
}
