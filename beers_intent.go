package main // import "github.com/Mumakil/toolo-skill"

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

const apiTimeFormat = "2006-01-02"
const incomingFormat = "2006-01-02"

// APIResponse is a response sent by the Töölö API
type APIResponse struct {
	CanHaveBeers bool `json:"can_have_beers"`
}

// BeersIntent handles the can have beers queries
type BeersIntent struct {
	APIURL string
}

// Handler handles the beers intent
func (bi *BeersIntent) Handler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	var date time.Time
	switch echoReq.GetIntentName() {
	case "GetToday":
		log.Println("Handling intent GetToday")
		date = time.Now()
	case "GetAnyDay":
		rawDate, err := echoReq.GetSlotValue("Date")
		if err != nil {
			log.Println("Date is missing:", err)
			echoResp.OutputSpeech("Sorry, I did not understand which date you’re talking about.")
			return
		}
		log.Println("Handling intent GetAnyDay with date", rawDate)
		parsedDate, err := time.Parse(incomingFormat, rawDate)
		if err != nil {
			log.Println("Error parsing query date:", err)
			echoResp.OutputSpeech("Sorry, I did not understand which date you’re talking about.")
			return
		}
		date = parsedDate
	}
	canHaveBeers, err := bi.fetchDataForDate(date)
	if err != nil {
		log.Println("Error fetching data from API:", err)
		echoResp.OutputSpeech("Sorry, I seem to have some problems fetching that information.")
		return
	}
	if canHaveBeers {
		echoResp.OutputSpeech("Yes, you can.")
	} else {
		echoResp.OutputSpeech("No, you can’t.")
	}
}

func (bi *BeersIntent) fetchDataForDate(date time.Time) (bool, error) {
	url := bi.APIURL + "/" + date.Format(apiTimeFormat)
	log.Println("GET", url)
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}
	var data APIResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	return data.CanHaveBeers, err
}
