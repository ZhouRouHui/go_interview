package main

import "fmt"

/**
2.下面代码输出什么，请说明？
*/
func two() {
	defer func() {
		fmt.Println(111)
		fmt.Println(recover())
	}()
	defer func() {
		fmt.Println(222)
		defer func() {
			fmt.Println(333)
			fmt.Println(recover())
		}()
		panic(1)
	}()
	defer recover()
	panic(2)

	/**
	参考答案及解析：1 2。
	recover() 必须在与 panic 同级的 defer() 函数中调用才有效，
	当程序中有多个层级的 panic 时，同级的 defer 函数里的 recover 优先处理这个 panic，

	引自《Go语言101》
	*/
}

func main() {
	two()
}
