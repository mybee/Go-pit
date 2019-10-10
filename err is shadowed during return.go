package main


import "fmt"
import "strconv"


func foo(x string) (ret int, err error) {
	ret = 1312

	if true {
		//var a int
		a, err1 := strconv.Atoi(x)
		if err1 != nil {
			err = err1
			return
		}
		fmt.Println(a)
	}
	return
}

func main() {
	fmt.Println(foo("123"))
}