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
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Summit struct {
	// SummitCode,AssociationName,RegionName,SummitName,AltM,AltFt,GridRef1,GridRef2,Longitude,Latitude,Points,BonusPoints,ValidFrom,ValidTo,ActivationCount,ActivationDate,ActivationCall
	SummitCode      string
	AssociationName string
	RegionName      string
	SummitName      string
	AltM            int
	AltFt           int
	GridRef1        string
	GridRef2        string
	Longitude       float64
	Latitude        float64
	Points          int
	BonusPoints     int
	ValidFrom       string
	ValidTo         string
	ActivationCount int
	ActivationDate  string
	ActivationCall  string
}

func parseLine(l []string) (s Summit, err error) {
	s.SummitCode = l[0]
	s.AssociationName = l[1]
	s.RegionName = l[2]
	s.SummitName = l[3]
	if l[4] != "" {
		s.AltM, err = strconv.Atoi(l[4])
	}
	if err != nil {
		return
	}
	if l[5] != "" {
		s.AltFt, err = strconv.Atoi(l[5])
	}
	if err != nil {
		return
	}

	s.GridRef1 = l[6]
	s.GridRef2 = l[7]

	if l[8] != "" {
		s.Longitude, err = strconv.ParseFloat(l[8], 32)
	}
	if err != nil {
		return
	}
	if l[9] != "" {
		s.Latitude, err = strconv.ParseFloat(l[9], 32)
	}
	if err != nil {
		return
	}
	if l[10] != "" {
		s.Points, err = strconv.Atoi(l[10])
	}
	if l[11] != "" {
		s.BonusPoints, err = strconv.Atoi(l[11])
	}
	//FIXME time.Time
	s.ValidFrom = l[12]
	s.ValidTo = l[13]

	if l[14] != "" {
		s.ActivationCount, err = strconv.Atoi(l[14])
	}

	//FIXME time.Time
	s.ActivationDate = l[15]

	s.ActivationCall = l[16]
	return
}

func ReadCSV(f *os.File) (summits []Summit, err error) {
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 17
	var l []string
	// ditch two header lines, ignore error on first line
	for i := 0; i < 2; i++ {
		l, _ = reader.Read()
	}
	var s Summit
	for i := 0; ; i++ {
		l, err = reader.Read()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
		s, err = parseLine(l)
		if err != nil {
			return
		}
		summits = append(summits, s)
	}
	return
}
