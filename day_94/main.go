package main

/**
1.下面这段代码输出什么？请简要说明。
*/
func one() {
	a := 2 ^ 15
	b := 4 ^ 15
	if a > b {
		println("a")
	} else {
		println("b")
	}

	/**
	参考答案及解析：a，^ 异或操作

	a: 0010 ^ 1111 = 1101 (2^15 = 13)
	b: 0100 ^ 1111 = 1011 (4^15 = 11)
	*/
}

/**
2.下面哪些函数不能通过编译？
*/
func A(string string) string {
	return string + string
}

func B(len int) int {
	return len + len
}

//func C(val, default string) string {
//	if val == "" {
//		return default
//	}
//	return val
//}

/**
参考答案及解析：C() 函数不能通过编译。C() 函数的 default 属于关键字。string 和 len 是预定义标识符，可以在局部使用。
nil 也可以当做变量使用，不过不建议写这样的代码，可读性不好，小心被接手你代码的人胖揍。

var nil = new(int)
func main() {
    var p *int
    if p == nil {
        fmt.Println("p is nil")
    } else {
        fmt.Println("p is not nil")
    }
}

*/

func main() {
	one()
}
