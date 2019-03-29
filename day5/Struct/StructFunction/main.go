package main

import "fmt"

type user struct {
	name string
	age int
}

//定义构造体内的方法
//aa 代表user struct实例
//(aa user) 代表是struct user 中的方法
//如果不使用指针,那么构造方法中的其他方法无法获取构造方法中的值
func (aa *user) structFunction(name string,age int)  {
	aa.name = name
	aa.age =age
	fmt.Println(name)
	fmt.Println(age)
}

func (aa user) get() user{
	return aa
}

func main()  {
	var stu user
	//调用构造体的方法
	stu.structFunction("username",22)
	//go会知道你想调用的是指针中的方法等同于((&stu))
	stu2 := stu.get()
	fmt.Println(stu2)
}
