package main

import (
	"fmt"
	"time"
)

func main() {
	data := []string{"one","two","three"}

	for _,v := range data {
		v := v
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}