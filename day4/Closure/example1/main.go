package main

import "fmt"

func main(){
	f:= add()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

func add() func(int) int{
	var x int
	return func(d int) int {
		//闭包保存当前x的值
		x +=d
		return x
	}
}

//res
//1
//101
//1101
