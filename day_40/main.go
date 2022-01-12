package main

import (
	"fmt"
)

/**
2.下面代码有什么问题？
*/
func one(stop <-chan int) {
	//close(stop)

	/**
	编译失败
	读类型通道不允许关闭，如果参数 stop 是 chan<- int 类型的话，就可以执行 close 操作
	*/
}

/**
3.下面这段代码存在什么问题？
*/

type Param map[string]interface{}

type Show struct {
	*Param
}

func two() {
	//s := new(Show)
	//s.Param["day"] = 2

	/**
	参考答案及解析：存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引。修复代码如下：

	func main() {
	    s := new(Show)
	    // 修复代码
	    p := make(Param)
	    p["day"] = 2
	    s.Param = &p
	    tmp := *s.Param
	    fmt.Println(tmp["day"])
	}
	*/
}

func main() {
	ch := make(chan int)
	one(ch)
	two()

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	select {
	case <-ch2:
		fmt.Println("ch2")
	case <-ch1:
		fmt.Println("ch1")
	}
}
