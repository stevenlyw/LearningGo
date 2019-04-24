package kernel

import (
	"LearningGo/go_pass_sdk/Constant"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type PassResponse struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Version int    `json:"version"`
}

type GetLicenseIdResponse struct {
	Code    int               `json:"code"`
	Data    map[string]string `json:"data"`
	Message string            `json:"message"`
	Version int               `json:"version"`
}

func GetLicenseId() (responseData map[string]interface{}) {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03@cn.primecloud.paas.windowsdemo"+Constant.USER_APPLY_LICENSE)
	request.Header.Set("Referer", "http://214.sgld.org")

	response, err := client.Do(request)
	responseData = make(map[string]interface{})

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

	passResponseData := &GetLicenseIdResponse{}
	err = json.Unmarshal(bytes, passResponseData)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "json解析失败"
		return
	}

	if passResponseData.Code != 200 {
		responseData["code"] = 500
		responseData["message"] = "状态码有误"
		return
	}

	responseData["code"] = 0
	responseData["data"] = passResponseData.Data["License"]
	return responseData
}


func DownloadUrl(fileId string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_DOWNLOAD_URL+"?FID="+fileId)
	request.Header.Set("Referer", "http://214.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}
	return passResponseDataHandle(response.Body)
}

func Transformation(fileId, convertType string, videoType string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)
	convertType = "0"
	videoType = "0"
	randomString := fmt.Sprintf("%x",time.Now().UnixNano())
	url := "?fileid="+fileId+"&ConverType="+convertType+"&VideoType="+videoType+"&random="+randomString
	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.CONVERT+url)
	request.Header.Set("Referer", "http://214.sgld.org")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	return passResponseDataHandle(response.Body)
}

func OnlineUrl(fileId,hostUrl string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)

	client := &http.Client{}

	url := "?FID="+fileId
	if hostUrl != "" {
		url = url+"&host"+hostUrl
	}
	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_ONLINEPLAY_URL+url)
	request.Header.Set("Referer", "http://214.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	return passResponseDataHandle(response.Body)
}

func GetFileInfo(fileId string,retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)

	client := &http.Client{}

	url := "?FID="+fileId

	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_FILE_INFO+url)
	request.Header.Set("Referer", "http://214.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	return passResponseDataHandle(response.Body)
}

func passResponseDataHandle(responseBody io.Reader) (responseData map[string]interface{}){
	bytes, err := ioutil.ReadAll(responseBody)
	responseData = make(map[string]interface{})
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	passResponseData := &PassResponse{}

	err = json.Unmarshal(bytes, passResponseData)

	switch passResponseData.Code {
	case 200:
		responseData["code"] = 0
		responseData["data"] = passResponseData.Data
		return
	case 404:
		responseData["code"] = 404
		responseData["message"] = "fileID不存在"
		return
	default:
		responseData["code"] = 500
		responseData["message"] = "获取下载地址失败"
		return
	}
}
