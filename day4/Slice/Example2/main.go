package main

import "fmt"

func main()  {
	arr := [...]int{1,2,3,4}
	slice := arr[:]
	fmt.Println(arr[1])
	fmt.Println(slice[1])
	//printf %p 是字符串(应该是...)
	fmt.Printf("slice[1]=%p",&slice[1])
	//append追加的方法
	slice = append(slice,5)
	//append批量添加的办法
	slice = append(slice, 4,5,6)
	fmt.Println(slice[1])
	fmt.Println(slice)
	sliceCopy()
	stringSlice()
}

func sliceCopy()  {
	//切片的拷贝
	var slice = []int{1,2,3,4,5}
	slice2 := make([]int,10)
	//如果初始化的切片比copy的切片要长,用0填充
	//[1 2 3 4 5 0 0 0 0 0]
	copy(slice2,slice)
	fmt.Println(slice2)
	slice3 := make([]int,4)
	//如果初始化的切片比copy的切片要短,那么就是从0开始截取
	//[1 2 3 4]
	copy(slice3,slice)
	fmt.Println(slice3)
}

func stringSlice(){
	//string的底层是一个byte数组,且是不可变的,也可以进行截取操作
	str := "hello world"
	slice1 := str[:5]
	slice2 := str[6:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	//如果要修改的话,放入一个新的切片中,然后修改
	//均为字符串可用byte
	//slice3 := []byte(str)
	//如果有汉字则需要使用rune防止乱码
	slice3 := []rune(str)
	slice3[0] = 'x'
	//转回string
	newString := string(slice3)
	fmt.Println(newString)
}
