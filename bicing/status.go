package bicing

type Status struct {
	UpdateTime int64     `json:"updateTime"`
	Stations   []Station `json:"stations"`
}
