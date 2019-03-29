package main

import "fmt"

type user struct {
	Name string
	Age int
}

func ConstructUser(name string, age int) *user {
	return &user{
		Name:name,
		Age:age,
	}
}


func main()  {
	 a := new(user)
	 a = ConstructUser("lalala", 50)
	 fmt.Println(a)
}
