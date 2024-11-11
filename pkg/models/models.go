package models

type CountryByIdRequest struct {
	Name string `json:"Name"`
}

type CountryModel struct {
	Name          string   `json:"Name"`
	TownsNames    []string `json:"TownsNames"`
	ContinentName string   `json:"ContinentName"`
}

type Town struct {
	Name  string `json:"Name"`
	Popul int    `json:"popul"`
}

type Continent struct {
	Name string `json:"Name"`
}
