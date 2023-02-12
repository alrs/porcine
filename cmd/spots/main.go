package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/alrs/sota"
	"github.com/davecgh/go-spew/spew"
)

const uri = "https://api2.sota.org.uk/api/spots/-1/all"

func fetchSpots() ([]sota.Spot, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var spots []sota.Spot
	err = json.Unmarshal(body, &spots)
	if err != nil {
		return nil, err
	}
	return spots, nil
}

func main() {
	seen := make(map[uint64]sota.Spot)
	firstRun := true
	for {
		log.Print("polling")
		spots, err := fetchSpots()
		if err != nil {
			log.Fatal(err)
		}
		sort.Slice(spots, func(i, j int) bool {
			return spots[i].ID < spots[j].ID
		})
		for _, s := range spots {
			_, present := seen[s.ID]
			if !present {
				if !firstRun {
					spew.Dump(s)
				}
				seen[s.ID] = s
			}
		}
		for k, v := range seen {
			d := time.Since(v.TimeStamp)
			if d > (24 * time.Hour) {
				log.Printf("forgetting %d from %s", k, v.TimeStamp)
				delete(seen, k)
			}
		}
		time.Sleep(30 * time.Second)
		firstRun = false
	}
}
