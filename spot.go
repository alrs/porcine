// package sota
// Copyright (C) 2023 Lars Lehtonen KJ6CBE

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
	"strconv"
	"strings"
	"time"

	"github.com/alrs/phonetic"
)

const timeFormat = "2006-01-02T15:04:05.999999999"

type Spot struct {
	ID                uint64    `json:"id"`
	UserID            uint64    `json:"userID"`
	TimeStamp         time.Time `json:"timeStamp"`
	Comments          string    `json:"comments"`
	Callsign          string    `json:"callsign"`
	AssociationCode   string    `json:"associationCode"`
	SummitCode        string    `json:"summitCode"`
	ActivatorCallsign string    `json:"activatorCallsign"`
	ActivatorName     string    `json:"activatorName"`
	Frequency         float32   `json:"frequency"`
	Mode              string    `json:"mode"`
	SummitDetails     string    `json:"summitDetails"`
	HighlightColor    string    `json:"highlightColor"`
}

func (s *Spot) Phonetic() string {
	code := phonetic.StringToNATO(
		fmt.Sprintf("%s/%s",
			s.AssociationCode,
			s.SummitCode),
	)
	freq := phonetic.StringToNATO(fmt.Sprintf("%.3f", s.Frequency))
	call := phonetic.StringToNATO(s.ActivatorCallsign)
	return fmt.Sprintf(
		"activation %s frequency %s mode %s callsign %s", code, freq, s.Mode, call,
	)
}

func (s *Spot) Summary() string {
	return fmt.Sprintf("%s/%s %.3f%s %s %q %s",
		s.AssociationCode,
		s.SummitCode,
		s.Frequency,
		s.Mode,
		s.ActivatorCallsign,
		s.ActivatorName,
		s.SummitDetails,
	)
}

func cleanFreq(freq string) string {
	freq = strings.TrimSpace(freq)
	freq = strings.ReplaceAll(freq, ":", ".")
	freq = strings.ReplaceAll(freq, ";", ".")
	freq = strings.ReplaceAll(freq, ",", ".")
	return freq
}

func (s *Spot) UnmarshalJSON(data []byte) error {
	type rawBlob struct {
		ID                uint64 `json:"id"`
		UserID            uint64 `json:"userID"`
		TimeStamp         string `json:"timeStamp"`
		Comments          string `json:"comments"`
		Callsign          string `json:"callsign"`
		AssociationCode   string `json:"associationCode"`
		SummitCode        string `json:"summitCode"`
		ActivatorCallsign string `json:"activatorCallsign"`
		ActivatorName     string `json:"activatorName"`
		Frequency         string `json:"frequency"`
		Mode              string `json:"mode"`
		SummitDetails     string `json:"summitDetails"`
		HighlightColor    string `json:"highlightColor"`
	}
	var rb rawBlob
	err := json.Unmarshal(data, &rb)
	if err != nil {
		return err
	}
	s.ID = rb.ID
	s.UserID = rb.UserID
	s.TimeStamp, err = time.Parse(timeFormat, rb.TimeStamp)
	if err != nil {
		return err
	}
	s.Comments = rb.Comments
	s.Callsign = rb.Callsign
	s.AssociationCode = rb.AssociationCode
	s.SummitCode = rb.SummitCode
	s.ActivatorCallsign = rb.ActivatorCallsign
	s.ActivatorName = rb.ActivatorName
	var f float64
	if rb.Frequency != "" {
		f, err = strconv.ParseFloat(cleanFreq(rb.Frequency), 32)
		if err != nil {
			return err
		}
	}
	s.Frequency = float32(f)
	s.Mode = strings.ToLower(rb.Mode)
	s.SummitDetails = rb.SummitDetails
	s.HighlightColor = rb.HighlightColor
	return nil
}
