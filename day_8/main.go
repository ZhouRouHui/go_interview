package main

import "fmt"

func hello() []string {
	return nil
}

/**
2.下面这段代码输出什么以及原因？
*/
func one() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}

	/**
	输出：not nil
	这道题目里面，是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。
	*/
}

/**
3.下面这段代码能否编译通过？如果可以，输出什么？
*/
func GetValue() int {
	return 1
}

func two() {
	//i := GetValue()
	//switch i.(type) {
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//case interface{}:
	//	fmt.Println("interface")
	//default:
	//	fmt.Println("unknown")
	//}

	/**
	参考答案及解析：编译失败。考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，
	需要注意的是，只有接口类型才可以使用类型断言和类型选择。
	*/
}

func main() {
	one()
	two()
}
