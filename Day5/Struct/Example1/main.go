package Example1

//定义构造体
type user struct {
	id int
	username string
	password string
}

func main()  {
	//初始化构造体
	var user1 user
	user1.id =22
	user1.username = "lalala"
	user1.password = "aaa"
	//初始化构造体方法2
	var user2 *user= &user{
		id:23,
		username:"balala",
		password:"wulala",
	}

	//初始化构造体方法3
	user3 := user{
		id:666,
		username:"wahaha",
		password:"666666",
	}
}
