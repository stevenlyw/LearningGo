package kernel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func DownloadUrl(fileId string)  {
	var licenseId = GetLicenseId()
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://209.sgld.org:10001", nil)
	request.Header.Set("User-Action", "PC://f286980ad1ae763bc5f0e0596959da03:"+licenseId+"@cn.primecloud.paas.windowsdemo/Resource/GetDownloadUrl")
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

	//LicenseId := responseData.Data["License"]
	//return LicenseId
}