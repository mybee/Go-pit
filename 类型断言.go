package main

import "fmt"

func main() {
	var data interface{} = "great"

	if res, ok := data.(string); ok {
		fmt.Println("[is an string] value =>", res)
	} else {
		fmt.Println("[not an string] value =>", data)
	}
}
