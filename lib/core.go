package lib

import (
	"encoding/base64"
	"net/http"
	"log"
	"github.com/spf13/viper"
)

var apiUrl string
var apiKey string

func initConfig() {
	siteName := viper.GetString("global.companyName")
	apiUrl = "https://" + siteName + ".teamwork.com/"
	apiKey = viper.GetString("global.apiKey")
}

func GetRequest(endPoint string) *http.Response {

	initConfig()

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl+endPoint, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(apiKey))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp

}


func basicAuth(apiKey string) string {
	return base64.StdEncoding.EncodeToString([]byte(apiKey))
}
