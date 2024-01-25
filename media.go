package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Media struct {
	ID        uint   `json:"id"`
	Guid      string `json:"guid"`
	MediaType string `json:"final_media_type_id"`
}
type Medias []Media

// Download picture using media guid
// ex: https://serviceapp.amisgest.ca/8_2/api/media/GetData?guid=54gA8qdDfcKYhyNcP66MDFtf1hsq7TUroE7K1fTwk2P2sW8iPPIfwfZr2XDjC5KaJu8xWnd7eSsGIohLrQJkhGRaWhghm1exIzht2LnUyQfta8QHPDHaDik4vlu1G5lj
func (m *Media) GetData(fileDate time.Time, destFolder string) error {
	if m.Guid == "" {
		return fmt.Errorf("media (id:%d) Guid is missing", m.ID)
	}
	res, err := http.Get(apiGetData + "?guid=" + m.Guid)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	filename := destFolder + "/" + fileDate.Format("2006-01-02") + "_" + strconv.Itoa(int(m.ID)) + "." + strings.Split(m.MediaType, "/")[1]
	fmt.Printf("Downloading %s\n", filename)
	err = os.WriteFile(filename, bytes, 0644)
	os.Chtimes(filename, fileDate, fileDate)
	if err != nil {
		return err
	}
	return nil
}

// Get media guid using the media ids
// ex: https://serviceapp.amisgest.ca/8_2/api/media/GetMediaGuid?mediaId=9067281
// func getMediaGuid(token Token, mediaId uint) Media {
func (m *Media) getGuid(token Token) {
	mediaIdStr := strconv.Itoa(int(m.ID)) // transform mediaId uint to string
	Get(
		apiGetMediaGuid+"?mediaId="+mediaIdStr,
		token,
		&m,
	)
}
