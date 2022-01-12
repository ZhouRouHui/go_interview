package main

import "fmt"

/**
1.下面这段代码输出什么？
*/

type Foo struct {
	bar string
}

func one() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo, len(s1))
	for i, value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0], s1[1], s1[2])
	fmt.Println(s2[0], s2[1], s2[2])

	/**
	参考答案及解析：
	{A} {B} {C}
	&{C} &{C} &{C}

	for range 使用短变量声明(:=)的形式迭代变量时，变量 i、value 在每次循环体中都会被重用，而不是重新声明。
	所以 s2 每次填充的都是临时变量 value 的地址，而在最后一次循环中，value 被赋值为{c}。因此，s2 输出的时候显示出了三个 &{c}。

	可行的解决办法如下：

	for i := range s1 {
	    s2[i] = &s1[i]
	}
	*/
}

/**
2.下面代码里的 counter 的输出值？
*/
func two() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)

	/**
	参考答案及解析：
	2 或 3

	map 是引用类型，for range 时复制也是原 map 的指针，所以删除会对原数据有效，另外 for range 一个 map 时 map 是无序的，
	当第一次迭代时不是 A 的话，就会吧 A 删除，此时 counter 就等于 2，当第一次迭代就是 A 时，counter 就等于 3.
	*/
}

func main() {
	one()
	two()
}
