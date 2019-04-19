package kernel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DownloadUrlResponse struct {
	Code int `json:"code"`
	Data string `json:"data"`
	Message string `json:"message"`
	Version int `json:"version"`
}

func DownloadUrl(fileId string,repeatTime int) (responseData map[string]interface{}){
	responseData = make(map[string]interface{})
	getLicenseData := GetLicenseId()
	if getLicenseData["code"] != 0{
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	aa:= getLicenseData["data"]
	var licenseId = fmt.Printf("%s",aa)
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo/Resource/GetDownloadUrl?FID="+fileId)
	request.Header.Set("Referer", "http://214.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	passResponseData := &DownloadUrlResponse{}

	err = json.Unmarshal(bytes,passResponseData)
	if passResponseData.Code == 404 {
		responseData["code"] = 404
		responseData["message"] = "获取下载地址失败"
	}
	responseData["code"] =0
	responseData["data"] = passResponseData.Data
	return
}