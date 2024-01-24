package main

import (
	"os"
	"time"
)

var daysToFetch time.Duration = 366 // number of past days to fetch. Starting today. Max is 366
var destFolder string = "./media"   // default destination folder for downloaded pictures

func main() {
	LoadEnv()
	destFolder := Getenv("AGP_DEST_FOLDER", destFolder)
	os.Mkdir(destFolder, 0755)
	credentials := Credentials{
		Username: Getenv("AGP_USERNAME", ""),
		Password: Getenv("AGP_PASSWORD", ""),
	}
	token := credentials.Login()
	token.myID = Getenv("AGP_MYID", "")
	logbook := Logbook{}
	timeframe := TimeFrame{}
	beginDateTime := time.Now().Add(time.Hour * 24 * -daysToFetch)
	endDateTime := time.Now()
	timeframe.Init(beginDateTime, endDateTime)
	logbook.Config(token)
	logbook.Search(timeframe)
	logbook.Infos()
	logbook.DownloadAllPictures()
}
