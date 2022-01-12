package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
1.下面代码有什么问题，请说明？
*/
func one() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))

	/**
	参考答案及解析：defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。

	修复代码：

	func main() {
	    f, err := os.Open("file")
	    if err != nil {
	        return
	    }
	    defer f.Close()

	    b, err := ioutil.ReadAll(f)
	    println(string(b))
	}
	*/
}

/**
2.下面代码输出什么，为什么？
*/
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v\n", r)
		}
	}()
	panic(1)
	panic(2)
}

func two() {
	f()

	/**
	参考答案及解析：recover:1。知识点：panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容。
	*/
}

func main() {
	one()
	two()
}
