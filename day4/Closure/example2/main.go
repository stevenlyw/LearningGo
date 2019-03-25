package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			println(suffix)
			return name + suffix
		}
		return name
	}
}

func main() {
	suffix1 := makeSuffix(".bmp")
	suffix2 := makeSuffix(".jpg")
	fmt.Println(suffix1("test"))
	fmt.Println(suffix2("test"))
}
