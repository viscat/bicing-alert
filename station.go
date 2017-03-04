package main

import "bicingalert/bicing"

type Station struct {
	Id           int
	Btype        string
	Lat          float64
	Long         float64
	StreetName   string
	StreetNumber string
	Altitude     int
	Nearby       string
}

func NewStation(station bicing.Station) Station {
	return Station{
		Id:           station.Id,
		Btype:        station.Btype,
		Lat:          station.Lat,
		Long:         station.Long,
		StreetName:   station.StreetName,
		StreetNumber: station.StreetNumber,
		Altitude:     station.Altitude,
		Nearby:       station.Nearby,
	}
}