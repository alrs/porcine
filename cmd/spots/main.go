package main

import (
	"log"
	"time"

	"github.com/alrs/porcine/sota"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	seen := make(map[uint64]sota.Spot)
	firstRun := true
	for {
		log.Print("polling")
		spots, err := sota.FetchSpots(-1)
		if err != nil {
			log.Fatal(err)
		}
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
