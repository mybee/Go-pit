package main

import "fmt"

func main() {


	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			// 没有标签的“break”声明只能从内部的switch/select代码块中跳出来
			// "goto"声明也可以完成这个功能。。。
			case false:
			default:

			goto loop
		}
	}
loop:
	fmt.Println("out!")
}
