package main

import "fmt"

/**
下面这段代码输出什么，说明原因。
*/

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for k, v := range slice {
		m[k] = &v
	}

	for k, v := range m {
		fmt.Println(k, "-->", *v)
	}
}

/*
0 -> 3
1 -> 3
2 -> 3
3 -> 3

参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，而不是元素的引用，
所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，
所有输出都是3.


正确写法
func main() {

     slice := []int{0,1,2,3}
     m := make(map[int]*int)

     for key,val := range slice {
         value := val	// 另外申请一个内存来存放当前的值
         m[key] = &value
     }

    for k,v := range m {
        fmt.Println(k,"===>",*v)
    }
}
*/
