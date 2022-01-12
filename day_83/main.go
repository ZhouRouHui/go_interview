package main

import "fmt"

/**
2.下面代码输出什么？
*/
type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func two() {
	d1 := data{"one"}
	d1.print()

	//var in printer = data{"two"}
	//in.print()

	/**
	参考答案及解析：编译报错。

	    cannot use data literal (type data) as type printer in assignment:
	    data does not implement printer (print method has pointer receiver)

	因为实现了 printer 接口的不是 data 的值类型，而是 data 的指针类型。
	修复代码：
	var in printer = &data{"two"}
	in.print()
	*/
}

func main() {
	two()
}
