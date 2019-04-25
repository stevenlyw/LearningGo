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

type Pass struct {
	Host string
	AppId string
	Port string
	Request http.Request
}

func Construct(host ,appId ,port string) *Pass{
	request, _ := http.NewRequest("GET", host+":"+port, nil)
	request.Header.Set("Referer", host)
	return &Pass{
		Host:host,
		AppId:appId,
		Port:port,
		Request:*request,
	}
}

//获取授权Id
func (c *Pass)GetLicenseId() (responseData map[string]interface{}) {
	client := &http.Client{}
	c.Request.Header.Set("User-Action", c.AppId+"@cn.primecloud.paas.windowsdemo"+Constant.USER_APPLY_LICENSE)
	response, err := client.Do(&c.Request)
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

//获取资源下载链接
func (c *Pass)DownloadUrl(fileId string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := c.GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)
	client := &http.Client{}

	request, _ := http.NewRequest("GET", c.Host+":"+c.Port, nil)
	request.Header.Set("User-Action", c.AppId+":"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_DOWNLOAD_URL+"?FID="+fileId)
	request.Header.Set("Referer", c.Host)

	response, err := client.Do(request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}
	return passResponseDataHandle(response.Body)
}

//资源转换
func (c *Pass)Transformation(fileId, convertType string, videoType string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := c.GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)
	randomString := fmt.Sprintf("%x",time.Now().UnixNano())
	url := "?fileid="+fileId+"&ConverType="+convertType+"&VideoType="+videoType+"&random="+randomString
	c.Request.Header.Set("User-Action", c.AppId+":"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.CONVERT+url)
	client := &http.Client{}
	response, err := client.Do(&c.Request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	return passResponseDataHandle(response.Body)
}

//获取在线播放地址
func (c *Pass)OnlineUrl(fileId,hostUrl string, retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := c.GetLicenseId()
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
	c.Request.Header.Set("User-Action", c.AppId+":"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_ONLINEPLAY_URL+url)
	response, err := client.Do(&c.Request)
	if err != nil {
		responseData["code"] = 500
		responseData["message"] = "信息读取失败"
		return
	}

	return passResponseDataHandle(response.Body)
}

//获取文件信息
func (c *Pass)GetFileInfo(fileId string,retry bool) (responseData map[string]interface{}) {
	responseData = make(map[string]interface{})
	getLicenseData := c.GetLicenseId()
	if getLicenseData["code"] != 0 {
		responseData["code"] = 500
		responseData["message"] = "获取licenseId失败"
		return
	}
	licenseId := getLicenseData["data"].(string)

	client := &http.Client{}

	url := "?FID="+fileId

	c.Request.Header.Set("User-Action", c.AppId+":"+licenseId+"@cn.primecloud.paas.windowsdemo"+Constant.GET_FILE_INFO+url)

	response, err := client.Do(&c.Request)
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
	default:
		responseData["code"] = passResponseData.Code
		responseData["message"] = passResponseData.Message
		return
	}
}
