package main

func main() {
/*		a := 11
		b := 22*/

	c := 10
	var d* int
	d = &c
	println(*d)

/*		mod(&a, &b)
		println(a, b)*/
/*		a := 123
		b := 123
		println(&a)
		println(&b)*/
}

func mod(a *int, b *int) {
	*a = 100
	*b = 10000
}
