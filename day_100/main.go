package main

import (
	"fmt"
	"sync"
)

/**
1.下面代码输出什么？
*/
func one() {
	m := map[string]int{
		"G": 7, "A": 1,
		"C": 3, "E": 5,
		"D": 4, "B": 2,
		"F": 6, "I": 9,
		"H": 8,
	}
	var order []string
	for k, _ := range m {
		order = append(order, k)
	}
	fmt.Println(order)

	/**
	参考答案即解析：按字母无序输出。知识点：遍历 map 是无序的。
	*/
}

/**
2.下面的代码有什么问题？
*/
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func two() {
	count := 1000
	wg := sync.WaitGroup{}
	wg.Add(count * 3)
	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		wg.Done()
	}
	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			u.Get(fmt.Sprintf("user_%d", i))
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")

	/**
	参考答案即解析：在执行 Get() 方法时可能报错。
	知识点：读写锁。虽然可以使用 sync.Mutex 做写锁，但是 map 是并发读写不安全的。map 属于引用类型，
	并发读写时多个协程是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系，
	会报错 “fatal error: concurrent map read and map write”。
	*/
}

func main() {
	one()
	two()
}
