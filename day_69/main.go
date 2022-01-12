package main

import "fmt"

/**
1.关于 slice 或 map 操作，下面正确的是？
*/
func one() {
	// a
	//var s []int
	//s = append(s, 1)

	// b
	//var m map[string]int
	//m["one"] = 1

	// c
	//var s []int
	//s = make([]int, 0)
	//s = append(s, 1)

	// d
	//var m map[string]int
	//m = make(map[string]int)
	//m["one"] = 1

	/**
	参考答案及解析：a c d
	*/
}

/**
2.下面代码输出什么？
*/
func test(x int) (func(), func()) {
	fmt.Printf("%p\n", &x)
	return func() {
			fmt.Println("a")
			fmt.Printf("%p\n", &x)
			println(x)
			x += 10
		}, func() {
			fmt.Println("b")
			fmt.Printf("%p\n", &x)
			println(x)
		}
}

func two() {
	a, b := test(100)
	a()
	b()

	/**
	参考答案及解析：100 110。知识点：闭包引用相同变量，通过没给地方的变量地址打印就可以看出来。
	*/
}

func main() {
	one()
	two()
}
