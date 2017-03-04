package bicing

type Station struct {
	Id           int     `json:"id,string"`
	Btype        string  `json:"type"`
	Lat          float64 `json:"latitude,string"`
	Long         float64 `json:"longitude,string"`
	StreetName   string  `json:"streetName"`
	StreetNumber string  `json:"streetNumber"`
	Altitude     int     `json:"altitude,string"`
	Slots        int     `json:"slots,string"`
	Bikes        int     `json:"bikes,string"`
	Nearby       string  `json:"nearby,string"`
	Status       string  `json:"status"`
}
