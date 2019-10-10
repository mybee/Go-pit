package main

import "fmt"

func doRecover() {
	// recover()函数可以用于获取/拦截panic
	 //prints: recovered => <nil>
}

func main() {
	defer func() {
		fmt.Println("recovered =>",recover())
		doRecover() //panic is not recovered
	}()

	panic("not good")
}