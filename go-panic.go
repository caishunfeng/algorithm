package main

import "fmt"

/**
reconver可以恢复当前goroutine的panic，在defer func() {}中才会生效
不对panic进行捕获，则会导致整个进程异常退出
*/
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main recover: ", err)
		}
	}()
	go runPanic()
	for {
	}
}

func runPanic() {
	defer fmt.Println("defer")
	// defer recover() 无效
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover：", err)
		}
		// panic("panic3")
	}()
	panic("panic")
	panic("panic2")
}
