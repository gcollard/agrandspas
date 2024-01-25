package main

import (
	"os"
	"strconv"
	"time"
)

func main() {

	loadEnv() // load environment variables

	// Set number of days to fetch
	var daysToFetch time.Duration // number of past days to fetch. Starting today. Max is 366
	daysToFetchInt, _ := strconv.Atoi(getenv("AGP_DAYS_TO_FETCH", "366"))
	if daysToFetchInt < 1 {
		daysToFetchInt = 1
	}
	if daysToFetchInt > 366 {
		daysToFetchInt = 366
	}
	daysToFetch = time.Duration(daysToFetchInt)

	// Create destination folder
	destFolder := getenv("AGP_DEST_FOLDER", "./Downloads")
	os.Mkdir(destFolder, 0755)

	credentials := Credentials{
		Username: getenv("AGP_USERNAME", ""),
		Password: getenv("AGP_PASSWORD", ""),
	}
	logbook := Logbook{
		Token:      credentials.Login(),
		DestFolder: destFolder,
	}
	timeframe := TimeFrame{}
	beginDateTime := time.Now().Add(time.Hour * 24 * -daysToFetch)
	endDateTime := time.Now()
	timeframe.init(beginDateTime, endDateTime)
	logbook.Search(timeframe)
	logbook.Infos()
	logbook.DownloadAllPictures()
}
