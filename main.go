package main // import "github.com/Mumakil/toolo-skill"

import (
	"flag"
	"log"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var appIDPtr = flag.String("appId", "", "Amazon Alexa Application ID for the skill")
var portPtr = flag.String("port", "3000", "Port to run web server in")
var apiURLPtr = flag.String("apiURL", "https://toolo-api.herokuapp.com/", "API to query the date information")

func main() {
	flag.Parse()

	appID := *appIDPtr
	if appID == "" {
		log.Fatal("Missing app id")
	}
	port := *portPtr

	beersIntent := &BeersIntent{
		APIURL: *apiURLPtr,
	}

	var Applications = map[string]interface{}{
		"/echo/beers": alexa.EchoApplication{
			AppID:    appID,
			OnIntent: beersIntent.Handler,
			OnLaunch: beersIntent.Handler,
		},
	}

	alexa.Run(Applications, port)
}
