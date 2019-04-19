package main

import (
	"fmt"
	"LearningGo/Day1/calculate/algorithm"
)

func main()  {
	plus :=algorithm.Plus(2,3)
	minus:=algorithm.Minus(2,3)

	fmt.Println("plus =",plus)
	fmt.Println("minus =",minus)
}
