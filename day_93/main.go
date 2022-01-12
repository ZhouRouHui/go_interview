package main

import "fmt"

/**
2.下面代码能通过编译吗？请简要说明。
*/
type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

func two() {
	//employee := new(Employee)
	//employee.SetName("Jack")

	/**
	参考答案及解析：编译不通过。当使用 type 声明一个新类型，它不会继承原有类型的方法集。
	*/
}

func main() {
	two()
}
