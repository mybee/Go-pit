package main

import "fmt"

func main() {

	// 在一个“nil”的slice中添加元素是没问题的
	//var s []int
	//fmt.Printf("%p\n", s)
	//s = append(s,1)
	//fmt.Printf("%p", s)
	//
	//// 但对一个map做同样的事将会生成一个运行时的panic
	//var m map[string] string
	//m = make(map[string]string)
	//m["one"] = "sfdfs"

	// 嵌套map
	var mm map[int]map[int]string
	mm = make(map[int]map[int]string)
	if _, ok := mm[1]; !ok {
		mm[1] = make(map[int]string)
		mm[1][2] = "21"
	}
	fmt.Println(mm)
}
