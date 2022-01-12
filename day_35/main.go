package main

import "fmt"

/**
定义 GetPodAction 的方法
*/

type Fragment interface {
	Exec() error
}

type GetPodAction struct {
}

func (g GetPodAction) Exec() error {
	return nil
}

func main() {
	// 初始化一个接口类型的方式
	var a Fragment = new(GetPodAction)
	var b Fragment = GetPodAction{}
	var c Fragment = &GetPodAction{}
	fmt.Println(a, b, c)
}
