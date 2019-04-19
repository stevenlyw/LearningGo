package main

import "fmt"

func main()  {
	FibonacciSequence(100)
}

func FibonacciSequence(n int){
	slice := make([]uint64, n)
	slice[0] = 1
	slice[1] = 1

	for i:=2;i<n ;i++  {
		slice[i] = slice[i-1] + slice[i-2]
	}
	fmt.Println(slice)
}
