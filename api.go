package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

/*
Data Model
----------------------------------------
LogbookEntry
	Activity
		ActivityMedia
			Media
	Observation
		ObservationMedia
			Media
*/

// frontend
var apiFrontendRoot = "https://app.journalapetitspas.ca/" // optional
var apiFrontendVersion = "28.3.5"                         // optional

// backend
var apiBackendRoot = "https://serviceapp.amisgest.ca/"
var apiBackendVersion = "8_2"
var apiBackend = apiBackendRoot + apiBackendVersion + "/"

// apiBackendAPI = apiBackend + "api/"
var apiBackendBreeze = apiBackend + "breeze/Breeze/"

// auth
var apiToken = apiBackend + "token"

// activities
var apiGetLogbookEntry = apiBackendBreeze + "logbook_entry"
var apiGetActivity = apiBackendBreeze + "activity"
var apiGetActivityMedia = apiBackendBreeze + "activity_media"

// observations
var apiGetObservations = apiBackendBreeze + "observation"
var apiGetObservationMedia = apiBackendBreeze + "observation_media"

// media
var apiBackendMedia = apiBackend + "api/media/"
var apiGetMediaGuid = apiBackendMedia + "GetMediaGuid" // full size
// apiGetMediaGuids = apiBackendMedia + "GetMedia"  // thumbnails
var apiGetData = apiBackendMedia + "GetData"

/* GET api endpoint wrapper */
func Get(url string, token Token, target interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// Set headers
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("authorization", "Bearer "+token.AccessToken)
	req.Header.Set("personid", token.myID)

	// Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	return json.Unmarshal(body, target)
}
