package main

/**
1.下面这段代码能通过编译吗？请简要说明。
*/
type T int

func F(t T) {}
func one() {
	//var q int
	//F(q)

	/**
	参考答案及解析：不能，int 和 T 是不同类型
	*/
}

/**
2.下面代码能通过编译吗？请简要说明。
*/
type T1 []int

func F1(t T1) {}

func two() {
	var q []int
	F1(q)

	/**
	参考答案及解析：能通过编译
	*/
}

func main() {
	one()
	two()
}
