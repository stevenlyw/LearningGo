package main

import (
	"fmt"
	"sort"
)

func main()  {
	sortInt()
	sortString()
	intSearch()
}

func sortInt(){
	//因为修改当前数组要传指针,所以要使用切片
	arr := []int{1,7,2,5,4,9,3}
	sort.Ints(arr)
	fmt.Println(arr)
}


func sortString(){
	arr := []string{"cbd","bcd","ace"}
	sort.Strings(arr)
	fmt.Println(arr)
}

func intSearch(){
	arr := []int{1,7,2,5,4,9,3}
	index := sort.SearchInts(arr[:], 3)
	fmt.Println(index)
}
