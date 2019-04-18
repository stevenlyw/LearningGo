package main

import (
	"LearningGo/go_pass_sdk/Kernel"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)
//import "LearningGo/go_pass_sdk/Kernel"

func main() {
	//fmt.Println(kernel.GetLicenseId())
	//kernel.DownloadUrl("b6f62afa966348038d0622cd656c66ce")
	var licenseId = kernel.GetLicenseId()
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://209.sgld.org:10001", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo/Resource/GetDownloadUrl?FID=b6f62afa966348038d0622cd656c66ce")
	request.Header.Set("Referer", "http://209.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}


	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
