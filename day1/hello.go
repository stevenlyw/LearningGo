package day1

import (
	"fmt"
	"LearningGo/day1/another"
	"LearningGo/day1/goroute"
	"time"
)

var pipe chan int
var c chan int

func main() {
	for i:= 0; i<100; i++ {
		go goroute(i)
		time.Sleep(10)
	}
	time.Sleep(10)
	sum ,avg := cal(1, 3)
	println("sum=", sum, "avg=", avg)

	sum := another.Add(2, 5)

	println(sum)

	c := make(chan int, 3)

	go bbb.Route(1,3,c)

	sum :=<- c


	println(sum)
	fmt.Println("|%b|", sum)
}

func add(a int, b int) {
	sum := a+ b
	pipe <- sum
}

func route(i int)  {
	println(i)
}

func cal (a int, b int) (int, int){
	sum := a+b
	avg := (a+b)/2

	return sum, avg
}
