package main

import (
	kernel "LearningGo/go_pass_sdk/Kernel"
	"fmt"
)

func main() {
	pass := kernel.Construct("http://214.sgld.org","PC://f286980ad1ae763bc5f0e0596959da03","13232")
	fmt.Println(pass.GetLicenseId())
	fmt.Println(pass.DownloadUrl("7c9c85b7e3964c1891626bede0ce138a",false))
	fmt.Println(pass.Transformation("7c9c85b7e3964c1891626bede0ce138a","0","1",false))
	fmt.Println(pass.OnlineUrl("7c9c85b7e3964c1891626bede0ce138a","",false))
	fmt.Println(pass.GetFileInfo("7c9c85b7e3964c1891626bede0ce138a",false))
}
