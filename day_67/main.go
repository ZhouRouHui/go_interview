package main

import "fmt"

type T struct {
	n int
}

/**
1.下面的代码输出什么？
*/
func one() {
	ts := [2]T{}
	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Println(ts[i].n, " ")
		}
	}
	fmt.Println(ts)

	/**
	参考答案及解析：9 [{0} {9}]。
	知识点：for-range 切片。for-range 切片时使用的是切片的副本，但不会复制底层数组，换句话说，此副本切片与原数组共享底层数组。
	*/
}

/**
2.下面的代码输出什么？
*/
func two() {
	fmt.Println()
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Println(t.n, " ")
		}
	}
	fmt.Println(ts)

	/**
	参考答案及解析：9 [{3} {9}]。知识点：for-range 切片。参考前几道题的解析，这道题的答案应该很明显。

	引自《Go语言101》
	*/
}

func main() {
	one()
	two()
}
