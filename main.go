package main

import (
	"os"
	"time"
)

var daysToFetch time.Duration = 366 // number of past days to fetch. Starting today. Max is 366

func main() {
	LoadEnv()

	credentials := Credentials{
		Username: Getenv("AGP_USERNAME", ""),
		Password: Getenv("AGP_PASSWORD", ""),
	}
	logbook := Logbook{
		Token:      credentials.Login(),
		DestFolder: Getenv("AGP_DEST_FOLDER", "./Downloads"),
	}
	os.Mkdir(logbook.DestFolder, 0755)
	timeframe := TimeFrame{}
	beginDateTime := time.Now().Add(time.Hour * 24 * -daysToFetch)
	endDateTime := time.Now()
	timeframe.Init(beginDateTime, endDateTime)
	logbook.Search(timeframe)
	logbook.Infos()
	logbook.DownloadAllPictures()
}
