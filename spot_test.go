package sota

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSpot(t *testing.T) {
	blob := `  {
    "id": 797214,
    "userID": 0,
    "timeStamp": "2023-02-12T18:00:08.54",
    "comments": "Javier [VK port-a-log]",
    "callsign": "EA2GM",
    "associationCode": "EA1",
    "summitCode": "AT-208",
    "activatorCallsign": "EA2GM/P",
    "activatorName": "Javier",
    "frequency": "21.040",
    "mode": "cw",
    "summitDetails": "Paisano, 636m, 2 pts",
    "highlightColor": "red"
  }`

	var s Spot

	err := json.Unmarshal([]byte(blob), &s)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(s)
}

func TestPhonetic(t *testing.T) {
	s := Spot{
		ID:                1234,
		AssociationCode:   "w6",
		SummitCode:        "5678-3",
		Mode:              "cw",
		Frequency:         146.520,
		ActivatorCallsign: "KJ6CBE",
	}
	t.Log(spew.Sdump(s.Phonetic()))

}
