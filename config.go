package main

type Config struct {
	APIUrl string
	APIkey string
}

type Response struct {
	Data []Day `json:"data"`
}

type Project struct {
	TotalSeconds int     `json:"total_seconds"`
	ProjectName  string  `json:"name"`
	Percent      float64 `json:"percent"`
}

type DailyStat struct {
	Date         string
	Projects     []Project
	TotalSeconds int
}

type Day struct {
	GrandTotal struct {
		TotalSeconds int `json:"total_seconds"`
		Hours        int `json:"hours"`
		Minutes      int `json:"minutes"`
		Seconds      int `json:"seconds"`
	} `json:"grand_total"`

	Projects []Project `json:"projects"`

	Range struct {
		Date string `json:"date"`
	} `json:"range"`
}
