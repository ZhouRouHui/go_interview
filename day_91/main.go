package main

/**
2.下面代码有什么问题？请指出。
*/
//func (m map[string]string) Set(key, value string) {
//	m[key] = value
//}

func two() {
	//m := make(map[string]string)
	//m.Set("a", "one")

	/**
	参考答案及解析：编译失败，Unnamed Type 不能作为方法的接收者。昨天我们讲过 Named Type 与 Unamed Type 的区别，就用 Named Type 来修复下代码：

	type User map[string]string

	func (m User) Set(key string, value string) {
	    m[key] = value
	}

	func main() {
	    m := make(User)
	    m.Set("A", "One")
	}
	*/
}

func main() {
	two()
}
