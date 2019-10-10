package main

import (
	"fmt"
)

type dataa struct {
	num int
	key *string
	items map[string]bool
}

func (this *dataa) pmethod() {
	fmt.Printf("%p  ", this)
	this.num = 7
	//this.key = "v.key"
	fmt.Printf("%p", this.items)
	this.items["vmethod"] = true
}

func (this dataa) vmethod() {
	fmt.Printf("%p  ", &this)
	this.num = 8
	*this.key = "v.key"
	fmt.Printf("%p", this.items)
	this.items["vmethod"] = false
}

func main() {
	key := "key.1"
	d := dataa{1,&key,make(map[string]bool, 3)}

	fmt.Printf("%p", &d)

	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=1 key=key.1 items=map[]

	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=7 key=key.1 items=map[]

	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	//prints num=7 key=v.key items=map[vmethod:true]
}