package main

//自定义的函数类型 customFunc
type customFunc func(int, int) int

func add(a int, b int) int {
	return a+b
}

func mining(a int, b int) int{
	return a-b
}

func operator(aaa customFunc, a int, b int) int {
	return aaa(a, b)
}

func main()  {
	//c :=add
	//所有符合该函数定义的都可以放入这个函数体
	//sum :=operator(c ,100, 200)
	sum := operator(mining, 100, 200)
	println(sum)
}
