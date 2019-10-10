package main

import (
	"fmt"
	"sync"
	"time"
)

type Demo struct {

	Data map[string]int

	Lock sync.Mutex

}



func (d *Demo) Get(k string) int{

	d.Lock.Lock()

	defer d.Lock.Unlock()

	return d.Data[k]

}



func (d *Demo) Set(k string, v int) {

	d.Lock.Lock()

	defer d.Lock.Unlock()

	d.Data[k]=v

}

func main(){

	//c := make(map[string]int)
	c := &Demo{Data:make(map[string]int, 0), Lock:sync.Mutex{}}
	go func() {//开一个协程写map
		for j := 0; j < 1000000; j++ {
			c.Set(fmt.Sprintf("%d", j), j)
			//c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {    //开一个协程读map
		for j := 0; j < 1000000; j++ {
			//fmt.Println(c[fmt.Sprintf("%d",j)])
			c.Get(fmt.Sprintf("%d",j))
		}
	}()

	time.Sleep(time.Second*20)

}
