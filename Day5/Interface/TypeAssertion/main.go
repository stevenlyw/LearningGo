package main

import "fmt"

func main()  {

	var a interface{}
	var b int
	a = b
	c:= a
	fmt.Printf("%T",c)
}
