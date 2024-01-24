package main

import (
	"time"

	"github.com/google/go-querystring/query"
)

var apiTimeFrameFormat = "2006-01-02T15:04:05.000Z"

type TimeFrame struct {
	BeginDateTime string `url:"beginDateTime"`
	EndDateTime   string `url:"endDateTime"`
	OwnerPersonID string `url:"ownerPersonId"`
}

// Init timeframe to look for with provided begin and end time
func (t *TimeFrame) Init(beginDateTime, endDateTime time.Time) {
	t.BeginDateTime = beginDateTime.Format(apiTimeFrameFormat)
	t.EndDateTime = endDateTime.Format(apiTimeFrameFormat)
	t.OwnerPersonID = "" // blank by default. Used to filter by person when multiple kids are in the same account
}

// Update timeframe begin datetime
func (t *TimeFrame) SetBeginDateTime(beginDateTime time.Time) {
	t.BeginDateTime = beginDateTime.Format(apiTimeFrameFormat)
}

// Update timeframe end datetime
func (t *TimeFrame) SetEndDateTime(endDateTime time.Time) {
	t.EndDateTime = endDateTime.Format(apiTimeFrameFormat)
}

/*
////////////////////////////////////////////
Observations by timeframe
observation -> get `id` and `compile_date`
- observation_media -> get `observation_id` and `media_id`
-- media : use `media_id` -> get `guid`
--- data : use `guid`
*/

/*
Get observation ids from timeframes
ex: https://serviceapp.amisgest.ca/8_2/breeze/Breeze/observation?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
*/
func (t *TimeFrame) getObservations(token Token) Observations {
	values, _ := query.Values(t)
	observations := Observations{}
	Get(
		apiGetObservations+"?"+values.Encode(),
		token,
		&observations,
	)
	return observations
}

/*
Get media ids using datetime frames
ex: https://serviceapp.amisgest.ca/8_2/breeze/Breeze/observation_media?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
*/
func (t *TimeFrame) getObservationMedias(token Token) ObservationMedias {
	values, _ := query.Values(t)
	ObservationMedias := ObservationMedias{}
	Get(
		apiGetObservationMedia+"?"+values.Encode(),
		token,
		&ObservationMedias,
	)
	return ObservationMedias
}

/*
////////////////////////////////////////////
Activities by timeframe
- logbook_entry -> get `posted_date` and `id`
- activity -> get `id` and `logbook_entry_id`
- activity_media -> get `media_id` and `activity_id`
-- media : use `media_id` -> get `guid`
--- data : use `guid`
*/

/*
Get logbook `posted_date` and `id` from timeframes
ex: https://serviceapp.amisgest.ca/8_2/breeze/Breeze/logbook_entry?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
*/
func (t *TimeFrame) getLogbookEntries(token Token) LogbookEntries {
	values, _ := query.Values(t)
	logbookEntries := LogbookEntries{}

	Get(
		apiGetLogbookEntry+"?"+values.Encode(),
		token,
		&logbookEntries,
	)
	return logbookEntries
}

/*
Get activities `id` and `logbook_entry` to reconcile with logbook_entry
this is a necessary step to get the `posted_date` from logbook_entry
ex: https://serviceapp.amisgest.ca/8_2/breeze/Breeze/logbook_entry?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
*/
func (t *TimeFrame) getActivities(token Token) Activities {
	values, _ := query.Values(t)
	activities := Activities{}
	Get(
		apiGetActivity+"?"+values.Encode(),
		token,
		&activities,
	)
	return activities
}

/*
Get activityMedia `media_id` and `activity_id` to reconcile with activity
this is a necessary step to get the `posted_date` from logbook_entry.activity
ex: https://serviceapp.amisgest.ca/8_2/breeze/Breeze/logbook_entry?beginDateTime=2024-01-23T05%3A00%3A00.000Z&endDateTime=2024-01-24T04%3A59%3A59.999Z&ownerPersonId=
*/
func (t *TimeFrame) getActivityMedia(token Token) ActivityMedias {
	values, _ := query.Values(t)
	activityMedias := ActivityMedias{}
	Get(
		apiGetActivityMedia+"?"+values.Encode(),
		token,
		&activityMedias,
	)
	return activityMedias
}
