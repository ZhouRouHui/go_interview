package main

import (
	"sync"
	"time"
)

/**
2.下面代码输出什么？请简要说明。
A. 不能编译；
B. 无输出，正常退出；
C. 程序 hang 住；
D. panic；
*/
func two() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()

	/**
	参考答案及解析：D。WaitGroup 在调用 Wait() 之后不能再调用 Add() 方法的。
	*/
}

func main() {
	two()
}
