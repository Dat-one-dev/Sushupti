package data

type Config struct {
	APIUrl string
	APIkey string
}

type Response struct {
	Data []Day `json:"data"`
}

type Project struct {
	TotalSeconds float64 `json:"total_seconds"`
	ProjectName  string  `json:"name"`
	Percent      float64 `json:"percent"`
}

type DailyStat struct {
	Date         string
	Projects     []Project
	Editors      []Category
	Languages    []Category
	TotalSeconds float64
}

type Category struct {
	Name         string  `json:"name"`
	TotalSeconds float64 `json:"total_seconds"`
	Percent      float64 `json:"percent"`
}

type Day struct {
	GrandTotal struct {
		TotalSeconds float64 `json:"total_seconds"`
		Hours        float64 `json:"hours"`
		Minutes      float64 `json:"minutes"`
		Seconds      float64 `json:"seconds"`
	} `json:"grand_total"`

	Editors   []Category `json:"editors"`
	Languages []Category `json:"languages"`
	Projects  []Project  `json:"projects"`

	Range struct {
		Date string `json:"date"`
	} `json:"range"`
}
