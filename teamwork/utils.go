package teamwork

import (
	"log"
	"os"

	"github.com/swill/teamwork"
	"github.com/rtslabs/teamwork-go/configuration"
)

var (
	conn *teamwork.Connection
)

// TeamworkConnection ...
func TeamworkConnection() *teamwork.Connection {

	twConfig, err := configuration.GetTeamworkConfig()
	if err != nil {
		log.Fatal(err)
	}

	// setup the teamwork connection
	conn, err := teamwork.Connect(twConfig.ApiKey)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}
