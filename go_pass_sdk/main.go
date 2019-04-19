package main

import (
	"LearningGo/go_pass_sdk/Kernel"
	"fmt"
)

type UploadStateResponse struct {
	Message string
	Data    map[string]interface{}
	Code    int
}

func main() {
	//var licenseId = kernel.GetLicenseId()
	result := kernel.DownloadUrl("e3a790bfe3734719881a443a293cfdab",0)

	//md5 := ""
	//filename := ""
	//client := &http.Client{}
	//
	//request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	//request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo/Resource/GetUploadState??md5="+md5+"&filename="+filename+"&directory=/")
	//request.Header.Set("Referer", "http://214.sgld.org")
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//bytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//responseData := &UploadStateResponse{}
	//
	//err = json.Unmarshal(bytes, responseData)
	//if responseData.Code == 404 {
	//	log.Fatal(responseData.Message)
	//}
	//fmt.Println(string(bytes))
}
