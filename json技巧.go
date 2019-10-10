package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type n struct {
	Name string `json:"name"`
}

type Server struct {
	// ID  不会导出到 JSON 中
	ID int `json:"-"`
	// ServerName 会在JSON中被 serverName 覆盖
	ServerName string `json:"serverName"`
	//转换成JSON字符串
	ServerName2 int32 `json:"serverName2" `
	//  如果 ServerIP  为空，则不输出到 JSON 串中
	ServerIP string `json:"serverIP,omitempty"`
	n
}

func main() {
	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: 22,
		ServerIP:    ``,
		n : n{Name:"231"},
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)


	var ss []Server
	str := `[{"serverName":"TianJin", "serverName2":22, "serverIP":"127.0.0.1"},
{"serverName":"Beijing", "serverName2":33, "serverIP":"127.0.0.2"}]`
	err := json.Unmarshal([]byte(str), &ss)
	if err != nil {
		panic(err)
	}
	fmt.Println(ss)
}

func mmain() {
	// 现在我们假设有如下的JSON数据
	b := []byte(`{"Name":"Wednesday","Age": 6.0, "Parents":["Gomez","Morticia"]}`)
	// 在我们不知道他的结构的情况下，我们把他解析到interface{}里面
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})
	// 通过断言之后，你就可以通过如下方式来访问里面的数据了
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int, int64, int8, int16, int32:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			//fmt.Println(reflect.TypeOf(v))
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
