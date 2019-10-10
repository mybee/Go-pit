package main

import "fmt"

func main() {
	//var i = 1
	//
	//// 被defer的函数的参数会在defer声明时求值, 参数值拷贝（而不是在函数实际执行时）。
	//defer fmt.Println("result =>",i*3)
	//defer fmt.Println("result =>",i)
	//defer fmt.Println("result =>",4)
	//i++
	//i = 100
	//prints: result => 2 (not ok if you expected 4)

	// 闭包 地址引用
	for j := 0; j< 3; j++ {
		defer fmt.Println(j)
		defer func() {
			fmt.Println("==>", j)
		}()
	}
}