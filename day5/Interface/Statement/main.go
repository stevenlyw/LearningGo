package Statement

import "fmt"

type student struct {
	id int
	name string
}

//interface 中只能有方法,不能有变量
type user interface {
	print()
}

func (aa *student) print()  {
	fmt.Println(aa.id)
	fmt.Println(aa.name)
}


func main() {
	var stu student
	var user1 user
	stu.id = 1
	stu.name = "lalala"
	//因为stu是指针类型实现的,所以如果接口想调用student构造体,必须也要使用指针
	user1 = &stu
	user1.print()
}
