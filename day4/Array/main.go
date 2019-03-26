package main

import (
	"fmt"
)

//数组:同一种数据类型,且长度固定
//实例 var a[length] type  var a[5] int
//长度是数组类型的一部分 a[5] int a[10] int 不是一种数据类型
//下标从零开始

func main()  {
	//一维数组初始化的几种方法
/*	var age = [5]int{1,2,3}
	var age = [...]int{1,2,3}
	//使用类型推导必须将数组填满
	age := [...]int{1,2,3,4,5}
	str := [3]string{0:"tony",1:"tom"}

	//多维数组初始化的几种方法
	var age [5][3]int
	var age [...][...]int*/
	var age = [...][3] int {{1,2,3}, {4,5,6}}
	fmt.Printf("%v",age)

	for _,val :=range age{
		fmt.Printf("%v\n", val)
		for _,value:= range val {
			fmt.Println(value)
		}
	}
}
