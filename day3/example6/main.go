package main

func main() {
	/*	i:= 10
		defer println(i)
		defer println("lalala")

		i = 100
		println(i)*/

/*	for i := 0; i < 5; i++ {
		defer println(i)
	}*/

/*	res := func(a int, b int) int {
		return a+b
	}(100,200)

	println(res)*/

	res := func(a int, b int) int {
		return a+b
	}

	println(res(100, 200))
}
