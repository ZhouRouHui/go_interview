package main

import "fmt"

/**
1.切片 a、b、c 的长度和容量分别是多少？
*/
func one() {
	s := [3]int{1, 2, 3}
	fmt.Println(len(s), cap(s))
	a := s[:0]         // 0, 3
	b := s[:2]         // 2, 3
	c := s[1:2:cap(s)] // 1, 2
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	fmt.Println(len(c), cap(c))
	/**
	参考答案及解析：
	0 3
	2 3
	1 2。
	知识点：数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组长度为 l。
	在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i。
	操作符 [i:j:k]，k 主要是用来限制切片的容量，但是不能大于数组的长度 l，截取得到的切片长度和容量计算方法是 j-i、k-i。
	*/
}

/**
2.下面代码输出什么？
*/
func increaseA() int {
	var i int
	fmt.Printf("%v\n", &i)
	defer func() {
		fmt.Printf("%v\n", &i)
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		fmt.Printf("%v\n", &r)
		r++
	}()
	fmt.Printf("%v\n", &r)
	return r
}

func two() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())

	/**
	参考答案解析
	0
	1
	知识点：defer、返回值。注意一下，increaseA() 的返回参数是匿名，increaseB() 是具名。
	*/
}

/**
3.下面这段代码输出什么？
*/
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func three() {
	var a A = Work{3}
	s := a.(Work)
	fmt.Println(s)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())

	/**
	参考答案及解析：13 23。
	知识点：类型断言
	*/
}

func main() {
	one()
	two()
	three()
}
