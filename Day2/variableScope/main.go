package main

import "fmt"

//在函数外部声明一个变量,作用域为全局,如果函数首字母大写外部也可引用
//不可使用 a:="G"方法
var a = "G"
func main()  {
	n()
	m()
	n()
}

func n()  {
	fmt.Println(a)
}

func m()  {
	//此处修改的是全局变量的值
	//如果为 a := "o",则为重新创建一个局部变量
	a = "o"
	fmt.Println(a)
}
