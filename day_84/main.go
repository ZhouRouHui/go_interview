package main

import "fmt"

/**
2.下面代码输出什么？
*/
func two() {
	a := [3]int{0, 1, 2}
	s := a[1:2]

	s[0] = 11         // 这里改变原数组
	s = append(s, 12) // 这里改变原数组
	s = append(s, 13) // 这里超出原数组长度了，所以重新分配了一个底层数组
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)

	/**
	输出：

	    [0 11 12]
	    [21 12 13]
	详情请参考《非懂不可的Slice（二）》
	*/
}

func main() {
	two()
}
