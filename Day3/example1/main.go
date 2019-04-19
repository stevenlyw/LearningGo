package main

import (
	"time"
)

func main() {
	/*	now := time.Now()
		a:= now.Format("2006-01-02 15:04:05")
		println(a)*/

/*	a := time.Now().UnixNano()
	cost()
	b := time.Now().UnixNano()

	fmt.Printf("cost:%d", (b-a)/1000)*/

	a := 11111
	b := 11112
	println(&a)
	println(&b)
}

func cost() {
	time.Sleep(time.Second * 5)
}
