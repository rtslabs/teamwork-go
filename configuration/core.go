package configuration

import (
	"encoding/base64"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var apiURL string
var apiKey string

func initConfig() {
	siteName := viper.GetString("global.companyName")
	apiURL = "https://" + siteName + ".teamwork.com/"
	apiKey = viper.GetString("global.apiKey")
}

// GetRequest ...
func GetRequest(endPoint string) *http.Response {

	initConfig()

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL+endPoint, nil)
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
