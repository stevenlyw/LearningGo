package main

import (
	"fmt"
	"math/rand"
	"time"
)

var a int

//在main函数前运行,修改rand随机数种子
func init()  {
	rand.Seed(time.Now().UnixNano())
}
func main()  {
   //res(5)
   //println(add.Name, add.Age)
   
   //a := 5
   //b := make(chan int, 3)
   //modify(&a)
   //
   //a:=10
   //b:=20
   //swap(&a,&b)
   //println(a, b)
   //
   //m()
   //println(a)

	for i:=0; i< 10; i++ {
		a:= rand.Int()
		println(a)
	}

	for i:=0;i<=10 ;i++  {
		a:=rand.Intn(100)
		println(a)
	}

	for i:=0;i<=10 ;i++  {
		a:=rand.Float32()
		println(a)
	}
}

func m()  {
	a = 100
}

func swap(a *int, b *int)  {
	tmp := *a
	*a = *b
	*b = tmp

}

func modify(a *int){
	*a = 10
	return
}

func res(n int){
	for i:=0; i<=n; i++{
		fmt.Printf("%d+%d=%d\n", i , n-i , n)
	}
}