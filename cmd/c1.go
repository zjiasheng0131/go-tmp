package main

import "fmt"

func main() {
	fmt.Println("程序开始运行...")

	// 创建一个 nil 指针
	var nilPointer *int

	// 故意访问 nil 指针，触发 panic
	fmt.Println("即将访问 nil 指针...")
	fmt.Println(*nilPointer) // 这里会崩溃

	fmt.Println("这行代码不会被执行")
}
