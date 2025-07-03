// Copyright (C) 2023-2025 Lars Lehtonen KJ6CBE

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
