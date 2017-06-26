package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// APIResponse is the Tarket wrapper for the json return
type APIResponse struct {
	IsDay bool     `json:"isDay"`
	Room  RoomJSON `json:"room"`
}

// RoomJSON is the struct reprensenting the tarket JSON structure
type RoomJSON struct {
	Department              string    `json:"department"`
	Room                    int       `json:"room"`
	RoomID                  string    `json:"roomId"`
	FallCount               int       `json:"fallCount"`
	ExitCount               int       `json:"exitCount"`
	IntrusionCount          int       `json:"intrusionCount"`
	GetUpCount              int       `json:"getUpCount"`
	BathroomCount           int       `json:"bathroomCount"`
	NightActivity           time.Time `json:"nightActivity"`
	NightActivityInBedroom  time.Time `json:"nightActivityInBedroom"`
	LastTimeFall            time.Time `json:"lastTimeFall"`
	LastTimeExit            time.Time `json:"lastTimeExit"`
	LastTimeIntrusion       time.Time `json:"lastTimeIntrusion"`
	LatestAlert             time.Time `json:"latestAlert"`
	DeviceIsDisabled        bool      `json:"deviceIsDisabled"`
	IntrusionAlarmActivated bool      `json:"intrusionAlarmActivated"`
	// Last event possible value :
	//('BEDROOM', 'BATHROOM', 'FALL', 'ABSENCE', 'PRESENCE', 'NO_SIGNAL')
	LastEvent string `json:"lastEvent"`
}

// TarketService is the generic interface for the remote services
type TarketService interface {
	GetRoomInfos(int) (APIResponse, error)
}

var currentTarketService TarketService

// GetTarketService return the singleton representing the tarket service
func GetTarketService() TarketService {
	if currentTarketService == nil {
		currentTarketService = newRemoteTarketService()
	}
	return currentTarketService
}

type remoteTarketService struct {
	sessionID string
	apiURL    string
	client    *http.Client
}

func newRemoteTarketService() remoteTarketService {

	r := remoteTarketService{}

	// We set some default value for convenience, feel free to change them
	r.apiURL = "http://front.recipe.fim-team.net/api/monitoring/room/FMDEV.500."
	r.sessionID = "588FCFD218D2A22EACC2EB0E6AC865D6"

	// We generate the http client that will execute our request
	// Let's do some reuse to save resources
	r.client = &http.Client{}
	return r
}

func (t remoteTarketService) GetRoomInfos(number int) (r APIResponse, err error) {

	// Create the request
	url := fmt.Sprint(t.apiURL, number, "?forLastDayPeriod=false")
	req, _ := http.NewRequest("GET", url, nil)

	// Generate the cookies, no need for AWSELB here. They just use JSESSIONID to authenticate
	cookie := http.Cookie{Name: "JSESSIONID", Value: t.sessionID}
	req.AddCookie(&cookie)

	// Execute and check for error
	resp, err := t.client.Do(req)
	if err != nil {
		log.Print(err)
	}

	// Don't forget to close the response body when this function return
	// or you will quickly get yourself caught in a "too many open file" error
	defer resp.Body.Close()

	// Decode the result of the request in our custom struct
	decodedResponse := APIResponse{}
	errorDecode := json.NewDecoder(resp.Body).Decode(&decodedResponse)

	return decodedResponse, errorDecode
}
