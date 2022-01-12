package main

import "fmt"

type person struct {
	name string
}

/**
1.通过指针变量 p 访问其成员变量 name，有哪几种方式？
A.p.name
B.(&p).name
C.(*p).name
D.p->name
参考答案及解析：AC。& 取址运算符，* 指针解引用。
*/
func one() {
	p := &person{
		name: "zrh",
	}
	fmt.Println(p.name)
	fmt.Println((*p).name)
}

/**
2.下面这段代码能否通过编译？如果通过，输出什么？
*/

type MyInt1 int
type MyInt2 = int

func two() {
	//var i int = 0
	//var i1 MyInt1 = i
	//var i2 MyInt2 = i
	//fmt.Println(i1, i2)
}

/**
这道题考的是类型别名与类型定义的区别。

MyInt1是基于类型 int 创建了新类型，MyInt2是创建了 int 的类型别名 ，注意类型别名的定义时 = 。所以，
当将 int 类型的变量赋值给 MyInt1 类型的变量时，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。

如果要将 int 类型的值赋值给 MyInt1 类型，可以使用强制类型转化 var i1 MyInt1 = MyInt1(i).
*/

func main() {
	one()
	two()
}
