package main // 1.定义包名

import ( // 2.包含的其他的包
	"fmt"
)

var a int // 3.定义全局变量和数据结构

func test(strMap map[int]string) bool { // 4.定义函数

	fmt.Println("Hello, playground: test!", strMap)
	for p := range strMap {
		strMap[p] = "Davis"
	}

	return true
}

func init() { // 5. main包中的init函数，会在main函数之前调用
	fmt.Println("Hello, playground: init begin!")
	a = 10
	fmt.Println("Hello, playground: init end", a)
}

func main() { // 6. 程序的入口main函数
	fmt.Println("Hello, playground: main begin!")
	strMap := make(map[int]string)
	strMap[0] = "hello"
	strMap[1] = "world"
	test(strMap)
	fmt.Println("Hello, playground: main end!", strMap)
}
