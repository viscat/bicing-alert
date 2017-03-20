package app

import (
	"bicingalert/bicing"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math"
	"time"
)

const MINUTES_NOT_UPDATED_WARNING = 5
const SECONDS_BETWEEN_UPDATES = 60

type Storage struct {
	Db mgo.Database
}

func (s Storage) UpdateBicingStatus() {

	lastUpdate := s.GetLastUpdate()

	for {
		status, err := bicing.GetStationsStatus()
		if err != nil {
			panic(err)
		}

		if lastUpdate != status.UpdateTime {
			log.Info("Updating status")
			status = s.update(status, lastUpdate)
		} else {
			log.Info("No new status update available")
		}

		if secondsFromLastUpdate := time.Now().Unix() - status.UpdateTime; secondsFromLastUpdate > MINUTES_NOT_UPDATED_WARNING*int64(time.Minute) {
			log.Warnf("Bicing API didn't update since %v seconds", secondsFromLastUpdate)

		}
		s.waitUntilNextUpdate(status)
	}

}

func (s Storage) GetLastUpdate() int64 {
	res := s.Db.C("StatusUpdate").Find(bson.M{}).Sort("-timestamp").Limit(1)
	statusUpdate := StatusUpdate{}
	res.One(&statusUpdate)
	lastUpdate := statusUpdate.Timestamp
	return lastUpdate
}

func (s Storage) waitUntilNextUpdate(status bicing.Status) {
	secondsForNextUpdate := s.getSecondsForNextUpdate(status)
	log.Infof("Next update will be in %v seconds", secondsForNextUpdate)
	time.Sleep(time.Duration(secondsForNextUpdate * int64(time.Second)))
}

func (s Storage) update(status bicing.Status, lastUpdate int64) bicing.Status {
	var stations []Station = make([]Station, len(status.Stations))
	var statusUpdate StatusUpdate = StatusUpdate{
		Timestamp:     status.UpdateTime,
		StationStatus: make([]StationStatus, len(status.Stations)),
	}

	for k, station := range status.Stations {
		stations[k] = NewStation(station)
		statusUpdate.StationStatus[k] = NewStationStatus(station)
		_, err := s.Db.C("Station").UpsertId(stations[k].Id, stations[k])
		if err != nil {
			log.Error(err)
		}
	}
	n, err := s.Db.C("StatusUpdate").Find(bson.M{"timestamp": status.UpdateTime}).Count()
	if n == 0 && err == nil {
		err = s.Db.C("StatusUpdate").Insert(statusUpdate)
		lastUpdate = statusUpdate.Timestamp
		if err != nil {
			log.Error(err)
		} else {
			log.Info("Bicing status updated correctly")
		}
	} else if err != nil {
		log.Error(err)
	}
	return status
}

func (s Storage) getSecondsForNextUpdate(status bicing.Status) int64 {
	return int64(math.Max(0, float64((status.UpdateTime+SECONDS_BETWEEN_UPDATES+1)-time.Now().Unix())))
}
