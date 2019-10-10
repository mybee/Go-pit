package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one","two","three"}

	for i,_ := range data {
		//go func(vv string) {
		data[i] = "four"
		//fmt.Println(v)
		//}(v)
	}

	fmt.Println(data)
	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}