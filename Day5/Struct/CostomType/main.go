package main

import "fmt"

//自定义一个新的类型 integer
type integer int

func main()  {
	var aa integer = 1000
	realInt := 100
	//自定义类型后,不能赋值给基础类型
	//realInt = aa
	//要使用类型转换方法
	realInt = int(aa)
	fmt.Println(realInt)
}
