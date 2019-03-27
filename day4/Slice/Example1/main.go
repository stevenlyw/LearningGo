package main

import "fmt"

//切片是数组的一种引用,在底层以数组形式表示,而且是引用类型
//切片是变长数组
//定义slice的方法  var slice = []int var slice = []int{1,2,3,4,5}
//定义后赋值slice的方法 1.创建数组后切片 2.make切片

func main()  {
	var slice []int
	//会在底层生成一个[0,0,0,0,0,]的数组
	//make创建
	//slice := make([]int,3, 5)

	//var slice2 = []int{1,2,3,4,5}
	arr := [5]int{1,2,3,4,5}
	//包含整个数组
	slice = arr[:]
	//从头开始,头不用写
	slice = arr[:5]
	//截取到尾,尾不用写
	slice = arr[1:]
	//切片截取 [start:end] 包含start不包含end 所以截取了3
	slice = arr[2:3]
	fmt.Println(slice) //[3]
	//查看切片长度len(),显示的是当前切片所含数据的个数 // 1
	fmt.Println(len(slice))
	//查看切片容量cap(),显示的初始生成的切片容量  // 3
	fmt.Println(cap(slice))
}
