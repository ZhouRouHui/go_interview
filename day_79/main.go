package main

import (
	"fmt"
)

/**
2.下面的代码有什么问题？
*/
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func two() {
	c := &ConfigOne{}
	c.String()
	/**
	参考答案及解析：无限递归循环，栈溢出。
	知识点：类型的 String() 方法。如果类型定义了 String() 方法，使用 Printf()、Print() 、 Println() 、 Sprintf() 等格式化输出时会自动使用 String() 方法。
	*/
}

func main() {
	two()
}
