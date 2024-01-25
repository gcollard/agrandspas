package main

import (
	"fmt"
	"time"
)

type Logbook struct {
	TimeFrame  TimeFrame
	Token      Token
	Entries    LogbookEntries
	DestFolder string
}

// https://serviceapp.amisgest.ca/8_2/breeze/Breeze/logbook_entry?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
type LogbookEntries []*LogbookEntry
type LogbookEntry struct {
	ID                   uint   `json:"id"`
	OwnerPersonID        uint   `json:"owner_person_id"`
	LogbookEntryStatusID uint   `json:"logbook_entry_status_id"`
	WriterPersonID       uint   `json:"writer_person_id"`
	CreationDate         string `json:"creation_date,omitempty"`
	ModificationDate     string `json:"modification_date,omitempty"`
	PostedDate           string `json:"posted_date,omitempty"`
	LocaleID             string `json:"locale_id"`
	Activities           Activities
	Observations         Observations
}

type Activities []*Activity
type Activity struct {
	ID             uint   `json:"id"`               // activity id
	LogbookEntryID uint   `json:"logbook_entry_id"` // coresponding logbook entry
	Order          uint   `json:"order"`            // frontend display order
	ActivityTypeID uint   `json:"activity_type_id"` // activity type id (ex: 1 for "meal", 2 for "diaper", 9 for "info")
	Comment        string `json:"comment"`          // activity comment from the teacher
	NumericValue   uint   `json:"numeric_value"`    // associated numeric value with the activity. ex: when activity_type = 1 (0="poor apetite", 3="very good apetite")
	ActivityMedias ActivityMedias
}

type ActivityMedias []*ActivityMedia
type ActivityMedia struct {
	MediaID    uint `json:"media_id"`    // media id
	ActivityID uint `json:"activity_id"` // activity id
	Media      Media
}

type Observations []*Observation
type Observation struct {
	ID                uint   `json:"id"`               // observation id
	LogbookEntryID    uint   `json:"logbook_entry_id"` // coresponding logbook entry
	Order             uint   `json:"order"`            // frontend display order
	Text              string `json:"text,omitempty"`   // comment
	IsCompiled        bool   `json:"is_compiled"`      // maybe: true if observation has been sent, false otherwise? (not sure) is it sometimes false?
	CompileDate       string `json:"compile_date"`     // datetime timestamp
	ObservationDate   string `json:"observation_date"` // datetime
	OwnerPersonID     uint   `json:"owner_person_id"`  // kid's ID
	WriterPersonID    uint   `json:"writer_person_id"` // group
	ObservationMedias ObservationMedias
}

type ObservationMedias []*ObservationMedia
type ObservationMedia struct {
	MediaID       uint `json:"media_id"`
	ObservationID uint `json:"observation_id"`
	Media         Media
}

// Search logbook entries by timeframe
func (l *Logbook) Search(timeframe TimeFrame) {
	l.TimeFrame = timeframe
	l.Entries = l.TimeFrame.getLogbookEntries(l.Token)
	l.getActivities()
	l.getObservations()
}

// Get all observations and observations medias for each logbook entry
func (l *Logbook) getObservations() {
	observations := l.TimeFrame.getObservations(l.Token)
	observationMedias := l.TimeFrame.getObservationMedias(l.Token)

	// merge observations and observationMedias
	for _, entry := range l.Entries {
		for _, observation := range observations {
			if observation.LogbookEntryID == entry.ID {
				for _, observationMedia := range observationMedias {
					if observation.ID == observationMedia.ObservationID {
						observationMedia.Media.ID = observationMedia.MediaID
						observationMedia.Media.getGuid(l.Token)
						observation.ObservationMedias = append(observation.ObservationMedias, observationMedia)
					}
				}
				entry.Observations = append(entry.Observations, observation)
			}
		}
	}
}

// Get all observations and observations medias for each logbook entry
func (l *Logbook) getActivities() {
	activities := l.TimeFrame.getActivities(l.Token)
	activitymedias := l.TimeFrame.getActivityMedia(l.Token)

	// merge activities and activityMedias
	for _, entry := range l.Entries {
		for _, activity := range activities {
			if activity.LogbookEntryID == entry.ID {
				for _, activityMedia := range activitymedias {
					if activity.ID == activityMedia.ActivityID {
						activityMedia.Media.ID = activityMedia.MediaID
						activityMedia.Media.getGuid(l.Token)
						activity.ActivityMedias = append(activity.ActivityMedias, activityMedia)
					}
				}
				entry.Activities = append(entry.Activities, activity)
			}
		}
	}
}

/* Download all pictures from their guids */
func (l *Logbook) DownloadAllPictures() {
	fmt.Printf("Downloading %d pictures\n", l.countPictures())
	l.downloadActivityPictures()
	l.downloadObservationsPictures()
}

// Download all observations attached pictures
func (l *Logbook) downloadObservationsPictures() {
	for _, entry := range l.Entries {
		for _, observation := range entry.Observations {
			for _, observationMedia := range observation.ObservationMedias {
				fileDate, _ := time.Parse(time.RFC3339, observation.CompileDate)
				// fmt.Println("downloading observation picture", fileDate, observation.Text, observationMedia.Media.ID)
				err := observationMedia.Media.GetData(fileDate, l.DestFolder)
				if err != nil {
					fmt.Println("error downloading observation picture", err)
				}
			}
		}
	}
}

// Download all activities attached pictures
func (l *Logbook) downloadActivityPictures() {
	for _, entry := range l.Entries {
		for _, activity := range entry.Activities {
			for _, activityMedia := range activity.ActivityMedias {
				fileDate, _ := time.Parse(time.RFC3339, entry.PostedDate)
				// fmt.Println("downloading activity picture", fileDate, activity.Comment, activityMedia.Media.ID)
				err := activityMedia.Media.GetData(fileDate, l.DestFolder)
				if err != nil {
					fmt.Println("error downloading activity picture", err)
				}
			}
		}
	}
}

/*
Return all infos about the logbook

1. Timeframe
2. Number of entries
3. Number of observations
4. Number of activities
5. Number of pictures
*/
func (l *Logbook) Infos() {
	fmt.Printf("Timeframe: %s - %s\n", l.TimeFrame.BeginDateTime, l.TimeFrame.EndDateTime)
	fmt.Printf("Number of entries: %d\n", len(l.Entries))
	fmt.Printf("Number of observations: %d\n", l.countObservations())
	fmt.Printf("Number of activities: %d\n", l.countActivities())
	fmt.Printf("Number of pictures: %d\n", l.countPictures())
}

func (l *Logbook) countObservations() int {
	count := 0
	for _, entry := range l.Entries {
		count += len(entry.Observations)
	}
	return count
}
func (l *Logbook) countActivities() int {
	count := 0
	for _, entry := range l.Entries {
		count += len(entry.Activities)
	}
	return count
}
func (l *Logbook) countActivityPictures() int {
	count := 0
	for _, entry := range l.Entries {
		for _, activity := range entry.Activities {
			count += len(activity.ActivityMedias)
		}
	}
	return count
}
func (l *Logbook) countObservationPictures() int {
	count := 0
	for _, entry := range l.Entries {
		for _, observation := range entry.Observations {
			count += len(observation.ObservationMedias)
		}
	}
	return count
}
func (l *Logbook) countPictures() int {
	return l.countActivityPictures() + l.countObservationPictures()
}
