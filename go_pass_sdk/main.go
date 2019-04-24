package main

func main() {

	//fmt.Println(a.(string))

	//newMd5 := md5.New()
	//newMd5.Sum([]byte(a))
	//fmt.Println(kernel.DownloadUrl("7c9c85b7e3964c1891626bede0ce138a",false))
	//fmt.Println(kernel.Transformation("7c9c85b7e3964c1891626bede0ce138a","0","0",false))
	//fmt.Println(kernel.OnlineUrl("7c9c85b7e3964c1891626bede0ce138a","",false))
	//fmt.Println(kernel.GetFileInfo("7c9c85b7e3964c1891626bede0ce138a",false))

	//responseData := make(map[string]interface{})
	//getLicenseData := kernel.GetLicenseId()
	//if getLicenseData["code"] != 0 {
	//	responseData["code"] = 500
	//	responseData["message"] = "获取licenseId失败"
	//	return
	//}
	//licenseId := getLicenseData["data"].(string)
	//fileId := "7c9c85b7e3964c1891626bede0ce138a"
	//convertType := "0"
	//videoType := "0"
	//newUUID, _ := uuid.NewV4()
	//random := newUUID.String()
	//client := &http.Client{}
	//
	//request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	//request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.CONVERT+"?fileid="+fileId+"&ConverType="+convertType+"&VideoType="+videoType+"&random="+random)
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
	////responseData := &UploadStateResponse{}
	////
	////err = json.Unmarshal(bytes, responseData)
	////if responseData.Code == 404 {
	////	log.Fatal(responseData.Message)
	////}
	//fmt.Println(string(bytes))
}
