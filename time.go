package main

import (
	"fmt"
	"time"
)

func main() {
	// 使用 2006-01-02 15:04:05 这样具有固定数字的格式化字符串
	t := time.Now().Format(time.RFC3339)
	fmt.Println(t)
}
