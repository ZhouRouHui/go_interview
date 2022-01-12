package main

import "fmt"

type T struct {
	n int
}

/**
1.下面的代码输出什么？
*/
func one() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			println(v)
		}()
		y[i] = &v
	}
	fmt.Println(y)
	println(*y[0], *y[1], *y[2])

	/**
	参考答案及解析：22222。
	知识点：defer()、for-range。
	for-range 虽然使用的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下。

	for-range 是开辟一个内存，每一次迭代就把新的内容覆盖这个内存，这里因为是取的这个内存的地址，所有所有打印结果都是 2
	*/
}

func main() {
	one()
}
