package main

import "fmt"

type carFactory struct {
	wheel int
}

func (car *carFactory) Run (){
	fmt.Println("running")
}

type bus struct {
	//通过匿名方式继承了carFactory,同时继承了变量和方法
	//匿名声明会引起歧义,不可重名
	carFactory
	name string
}

type train struct {
	//通过组合方式继承了carFactory,同时继承了变量和方法
	carFactory carFactory
}

func main() {
	var car bus
	car.wheel = 4
	car.name = "aaaa"
	car.Run()
	fmt.Println(car)

	var train train
	//调用时加定义的名称即可
	train.carFactory.Run()
}
