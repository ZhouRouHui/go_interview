package main

import "fmt"

/**
1.下面列举的是 recover() 的几种调用方式，哪些是正确的？
*/

func one() {
	// a
	//recover()
	//panic(1)

	// b
	//defer recover()
	//panic(1)

	// c
	//defer func() {
	//	recover()
	//}()
	//panic(1)

	// d
	//defer func() {
	//	defer func() {
	//		recover()
	//	}()
	//}()
	//panic(1)

	/**
	参考答案及解析：C。
	recover() 必须在 defer() 函数中直接调用才有效。
	上面其他几种情况调用都是无效的：直接调用 recover()、在 defer() 中调用 defer 再调用 recover() 和 defer() 调用时多层嵌套。
	*/
}

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
		defer fmt.Println(recover())
		panic(1)
	}()
	defer recover()
	panic(2)

	/**
	参考答案及解析：21。
	recover() 必须在与 panic 同级的 defer() 函数中调用才有效，
	当程序中有多个层级的 panic 时，同级的 defer 函数里的 recover 优先处理这个 panic，

	所以 defer recover() 是无效的。
	在调用 defer() 时，便会计算函数的参数并压入栈中，
	所以执行 defer fmt.Print(recover()) 时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。
	所以输出 21。

	引自《Go语言101》
	*/
}

func main() {
	one()
	two()
}
