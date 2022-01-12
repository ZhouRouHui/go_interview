package main

import "fmt"

/**
2.下面这段代码输出什么？
*/
func one() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nill")

	/**
	参考答案及解析：nil。当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
	*/
}

/**
3.下面这段代码输出什么？
*/
func two() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])

	/**
	参考答案及解析：0。删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。
	*/
}

func main() {
	one()
	two()
}
