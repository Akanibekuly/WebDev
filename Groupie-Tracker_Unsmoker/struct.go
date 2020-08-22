package main

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

type Artist struct {
	ID            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"Members"`
	CreationDates int      `json:"creationDates"`
	FirstAlbum    string   `json:"firstAlbum"`
	ConcertDates  string   `json:"concertDates"`
	Locations     string   `json:"locations"`
	Relations     string   `json:"relations"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
