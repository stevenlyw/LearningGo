package kernel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type PaasResponse struct {
	Code    int               `json:"code"`
	Data    map[string]string `json:"data"`
	Message string            `json:"message"`
	Version int               `json:"version"`
}

func GetLicenseId() string {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://209.sgld.org:10001", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03@cn.primecloud.paas.windowsdemo/user/UserApplyLicense")
	request.Header.Set("Referer", "http://209.sgld.org")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseData := &PaasResponse{}
	err = json.Unmarshal(bytes, responseData)
	if err != nil {
		log.Fatal(err)
	}
	if responseData.Code != 200 {
		log.Fatal("状态码有误")
	}

	LicenseId := responseData.Data["License"]
	return LicenseId
}


