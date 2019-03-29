package main

func main() {
	sum :=add(1, 2, 3, 4)
	println(sum)
}

func add(a int, arg ... int) (sum int) {
	sum = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return
}
