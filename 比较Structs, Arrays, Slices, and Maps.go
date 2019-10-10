package main

import "fmt"

type ddata struct {
	num int                //ok
	checks [10]func() bool //not comparable
	doit func() bool       //not comparable
	m map[string] string   //not comparable
	bytes []byte           //not comparable
}

func main() {
	v1 := ddata{}
	v2 := ddata{}
	fmt.Println("v1 == v2:",v1 == v2)
}