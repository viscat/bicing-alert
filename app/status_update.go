package app

import "bicingalert/bicing"

type StatusUpdate struct {
	Timestamp     int64
	StationStatus []StationStatus
}

type StationStatus struct {
	StationId int
	Slots     int
	Bikes     int
	Status    string
}

func NewStationStatus(station bicing.Station) StationStatus {
	return StationStatus{
		StationId: station.Id,
		Bikes:     station.Bikes,
		Slots:     station.Slots,
		Status:    station.Status,
	}
}
