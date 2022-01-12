package main

import "fmt"

func main() {
	one()
	two()
	three()
}

/**
1.下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。
*/
func one() {
	list := new([]int)
	fmt.Printf("%T, %v", list, list) // *[]int, &[]
	//list = append(list, 1)	编译失败
	//fmt.Println(list)

	/*
		参考答案及解析：不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。
		可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
	*/
}

/**
2.下面这段代码能否通过编译，如果可以，输出什么？
*/
func two() {
	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2)	// 编译失败，应该这么写 s1 = append(s1, s2...)
	//fmt.Println(s1)

	/**
	参考答案及解析：不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。
	或者直接跟上元素，形如：append(s1,1,2,3)。
	*/
}

var (
//size := 1024
//max_size = size * 2
)

/**
3.下面这段代码能否通过编译，如果可以，输出什么？
*/
func three() {
	//fmt.Println(size, max_size)

	/**
	编译失败
	全局变量不能使用 := 来申明和赋值
	*/
}
