package main

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	m["x"] = data{name:"two"} //error
}
