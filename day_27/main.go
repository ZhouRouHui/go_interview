package main

import "fmt"

/**
1.下面这段代码输出什么？为什么？
*/

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// 实现 Stringer 接口
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

func one() {
	fmt.Println(South)

	/**
	参考答案及解析：South。知识点：iota 的用法、类型的 String() 方法。

	根据 iota 的用法推断出 South 的值是 2；另外，如果类型定义了 String() 方法，
	当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印。
	*/
}

/**
2.下面代码输出什么？
*/
type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func two() {
	//m["foo"].x = 4
	//fmt.Println(m["foo"])

	/**
	参考答案及解析: 编译报错 cannot assign to struct field m[“foo”].x in map。
	错误原因：对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。

	有两个解决办法：

	1.使用临时变量
	type Math struct {
	    x, y int
	}

	var m = map[string]Math{
	    "foo": Math{2, 3},
	}

	func main() {
	    tmp := m["foo"]
	    tmp.x = 4
	    m["foo"] = tmp
	    fmt.Println(m["foo"].x)
	}

	2.修改数据结构
	type Math struct {
	    x, y int
	}

	var m = map[string]*Math{
	    "foo": &Math{2, 3},
	}

	func main() {
	    m["foo"].x = 4
	    fmt.Println(m["foo"].x)
	    fmt.Printf("%#v", m["foo"])   // %#v 格式化输出详细信息
	}
	*/
}

func main() {
	one()
	two()
}
