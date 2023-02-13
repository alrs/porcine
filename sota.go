package sota

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

const uriTemplate = "https://api2.sota.org.uk/api/spots/%d/all"

func FetchSpots(offset int) ([]Spot, error) {
	uri := fmt.Sprintf(uriTemplate, offset)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var spots []Spot
	err = json.Unmarshal(body, &spots)
	if err != nil {
		return nil, err
	}
	sort.Slice(spots, func(i, j int) bool {
		return spots[i].ID < spots[j].ID
	})

	return spots, nil
}
