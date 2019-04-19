package io

import (
	"log"
	"os"
)

func main() {
	file,err :=os.OpenFile("C:/aaa.log",os.O_RDWR|os.O_CREATE,0666)
	if err !=nil {
		log.Fatal(err)
	}


}
