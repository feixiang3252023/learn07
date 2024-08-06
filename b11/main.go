package main

import "fmt"

func main() {
	var i interface{} = "hello,world"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
