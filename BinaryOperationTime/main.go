package main

import (
	"fmt"
	"time"
)

func main() {
	arr := [10000000][]byte{}
	begin := time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		arr[i] = []byte("你好吗")
	}
	end := time.Now().UnixNano()
	aaTime := end - begin
	fmt.Println(aaTime/1e6)

	begin = time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		arr[i] = []byte{228, 189, 160, 229, 165, 189, 229, 144, 151}
	}
	end = time.Now().UnixNano()
	bbTime := end - begin
	fmt.Println(bbTime/1e6)

	begin =time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		arr[i] = []byte{0xe4,0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0xe5, 0x90, 0x97}
	}
	end = time.Now().UnixNano()
	ccTime := end-begin
	fmt.Println(ccTime/1e6)
}
