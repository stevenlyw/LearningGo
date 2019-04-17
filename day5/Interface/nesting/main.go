package main

import "fmt"

type interface1 interface {
	aaa()
	bbb()
}

type interface2 interface {
	ccc()
}

//interface3 嵌套了 前两个接口的所有方法
type interface3 interface {
	interface1
	interface2

	//ccc() 为避免错误不允许声明相同名字的方法,就算返回值类型不同也不可以
}

type struct1 struct {

}

func (struct1) aaa() {
	fmt.Println("我是方法aaa")
}

func (struct1) bbb() {
	fmt.Println("我是方法bbb")
}

func (struct1) ccc() {
	fmt.Println("我是方法ccc")
}

func main() {
	aa:= new(struct1)
	//因为struct1实现了interface3的所有方法,所以可以这样调用
	Test(aa)
}

func Test(rw interface3)  {
	rw.aaa()
	rw.bbb()
}


