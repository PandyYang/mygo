package main

import "fmt"

// 闭包函数示例
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 创建一个 add 函数，sum 初始值为 0
	add := adder()

	// 对 add 函数进行调用
	fmt.Println(add(1)) // 1
	fmt.Println(add(2)) // 3
	fmt.Println(add(3)) // 6

	// 创建一个新的 add 函数，sum 初始值也为 0
	addNew := adder()
	fmt.Println(addNew(1)) // 1
	fmt.Println(addNew(2)) // 3
}
