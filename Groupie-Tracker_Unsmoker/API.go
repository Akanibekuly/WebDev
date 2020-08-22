package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Gruopie struct {
	Artists   []Artist
	Dates     []Dates
	Locations []Locations
	Relations []Relations
}

func (g *Gruopie) getArtists() error {
	res, err := http.Get(artistsURL)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &g.Artists)
	return nil
}

func (g *Gruopie) getArtistById(n int) Artist {
	var artist Artist
	if len(g.Artists) < 1 || n > len(g.Artists) {
		return artist
	}
	return g.Artists[n-1]
}
