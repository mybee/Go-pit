package main

import "fmt"

func main() {
	data := []int{1,2,3}
	i := 0
	i++
	++i //error
	fmt.Println(data[i++]) //error
}