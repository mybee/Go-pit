package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	workerCount := 5
//
//	for i := 0; i < workerCount; i++ {
//		go doit(i)
//	}
//	time.Sleep(1 * time.Second)
//	fmt.Println("all done!")
//}
//
//func doit(workerId int) {
//	fmt.Printf("[%v] is running\n",workerId)
//	time.Sleep(3 * time.Second)
//	fmt.Printf("[%v] is done\n",workerId)
//}



func main() {
	var wg sync.WaitGroup

	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[%v] is running\n",workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n",workerId)
}