package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var input int
	number := rand.Intn(100)

	for {
		fmt.Scanf("%d\n", &input)
		flags := false
		switch {
		case input < number:
			println("to small")
		case input == number:
			flags =true
			println("right")
		case input > number:
			println("too big")
		}

		if flags {
			break
		}
	}
}
