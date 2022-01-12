package main

import "fmt"

/**
2. 下面代码输出什么？
*/
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
	//return "this is a pointer of ConfigOne"
}

func one() {
	//c := &ConfigOne{}
	//c.String()

	/**
	参考答案及解析：运行时错误。

	如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。上面这段代码是在该类型的 String() 方法内使用格式化输出，
	导致递归调用，最后抛错。
	*/
}

/**
3.下面代码输出什么？
*/
func two() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)
	fmt.Println(a, len(a), cap(a), &a)
	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
			fmt.Println(a, len(a), cap(a), &a) //
		}
		r = append(r, v)
	}
	fmt.Println(a)
	fmt.Println(r)

	/**
	参考答案及解析：[1 2 3 4 5]。
	a 在 for range 过程中增加了两个元素，len 由 5 增加到 7，但 for range 时会使用 a 的副本 a’ 参与循环，副本的 len 依旧是 5，
	因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素。
	 */
}

func main() {
	one()
	two()
}
