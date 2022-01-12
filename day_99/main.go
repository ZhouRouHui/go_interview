package main

import "fmt"

/**
1.下面这段代码能通过编译吗？请简要说明。
*/
func one() {
	true := false
	fmt.Println(true)

	/**
	参考答案及解析：通过编译，true 是预定义标识符可以用作变量名，但是不建议这么做。
	*/
}

/**
2.下面的代码输出什么？
*/
func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		fmt.Printf("%p\n", &ret) //0xc00001e0a0	这才是真正返回的 ret
		ret := 10
		fmt.Printf("%p\n", &ret) //0xc00001e0a8
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func two() {
	result := watShadowDefer(50)
	fmt.Println(result)

	/**
	参考答案即解析：100。知识点：变量作用域和defer 返回值。
	*/
}

func main() {
	one()
	two()
}
