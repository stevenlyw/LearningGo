package main

import "fmt"

func main()  {
	//map是kv的数据结构,又称字典或关联数组
	//声明map, 声明是不会分配内存的,需要初始化
	//var a map[keytype]valuetype
	//var a map[string]string
	//var a map [int]string
	//var a map [int]map[string]int
	//然后初始化
	//a =make(map[string]string, 10)
	//或者直接赋值
	//a :=make(map[string]string, 10)
	//或者直接声明
	var a = map[string]string{
		"abc":"value",
	}
	a ["def"] = "efg"
	//直接声明后,声明中的值有可能会排序在赋值语句值的后面(原因不明...)
	fmt.Println(a)
	//mapmap()
}

func mapmap()  {
	a := make(map[int]map[string]string)
	a[1] = make(map[string]string,10)
	a[1]["aaa"]= "bbb"

	//查看键是否存在,存在赋值给bb变量,值为val
	val,bb :=a[2]
	if bb {
		fmt.Println(val)
	}else {
		fmt.Println("error")
	}
	fmt.Println(a)
}
