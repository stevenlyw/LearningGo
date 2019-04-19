package main

import "os"

func main()  {
	file := os.Open(*filename)
	//defer 关键字 是在函数执行的最后执行该语句,如果存在多个,将以栈的顺序执行
	//defer用途,关闭文件,锁的释放,数据库链接释放
	defer file.Close()
}
