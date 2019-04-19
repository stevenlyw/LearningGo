package kernel

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetLicenseIdResponse struct {
	Code    int               `json:"code"`
	Data    map[string]string `json:"data"`
	Message string            `json:"message"`
	Version int               `json:"version"`
}

func GetLicenseId() (responseData map[string]interface{}) {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://214.sgld.org:13232", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03@cn.primecloud.paas.windowsdemo/user/UserApplyLicense")
	request.Header.Set("Referer", "http://214.sgld.org")

	responseData = make(map[string]interface{})
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

	passResponseData := &GetLicenseIdResponse{}
	err = json.Unmarshal(bytes, responseData)
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


