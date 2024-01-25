package main

import (
	"os"
	"time"
)

var daysToFetch time.Duration = 366 // number of past days to fetch. Starting today. Max is 366

func main() {
	loadEnv()
	credentials := Credentials{
		Username: getenv("AGP_USERNAME", ""),
		Password: getenv("AGP_PASSWORD", ""),
	}
	logbook := Logbook{
		Token:      credentials.Login(),
		DestFolder: getenv("AGP_DEST_FOLDER", "./Downloads"),
	}
	os.Mkdir(logbook.DestFolder, 0755)
	timeframe := TimeFrame{}
	beginDateTime := time.Now().Add(time.Hour * 24 * -daysToFetch)
	endDateTime := time.Now()
	timeframe.init(beginDateTime, endDateTime)
	logbook.Search(timeframe)
	logbook.Infos()
	logbook.DownloadAllPictures()
}
