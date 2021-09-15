package main

import "fmt"

/**
defer为在函数退出时执行，顺序为后进先出；
*/
func main() {
	for i := 0; i < 5; i++ {
		// defer结构体保留的是编译后的语句，参数已被保留
		defer fmt.Println("one:", i)
	}
	for i := 0; i < 5; i++ {
		// defer结构体保留的是函数指针，在最后执行时才获取参数
		defer func() {
			fmt.Println("two:", i)
		}()
	}
}
