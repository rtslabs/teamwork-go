package teamwork

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/swill/teamwork"
)

var (
	conn       *teamwork.Connection
)


// TeamworkConnection ...
func TeamworkConnection() *teamwork.Connection {
	apiToken := viper.GetString("global.apiKey")

	// setup the teamwork connection
	conn, err := teamwork.Connect(apiToken)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}
