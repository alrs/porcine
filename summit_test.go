package sota

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestReadCSV(t *testing.T) {
	f, err := os.Open("./testlist.csv")
	if err != nil {
		t.Fatal(err)
	}

	summits, err := ReadCSV(f)
	if err != nil {
		t.Fatal(err)
	}

	want := Summit{
		SummitCode:      "4O/IC-003",
		AssociationName: "Montenegro",
		RegionName:      "Istok Crne Gore",
		SummitName:      "Veliki vrh (Maja Gurt e Zjarmit)",
		AltM:            2480,
		AltFt:           8136,
		GridRef1:        "19.7872",
		GridRef2:        "42.4971",
		Longitude:       19.787200927734375,
		Latitude:        42.497100830078125,
		Points:          10,
		BonusPoints:     3,
		ValidFrom:       "01/03/2019",
		ValidTo:         "31/12/2099",
		ActivationCount: 0,
		ActivationDate:  "",
		ActivationCall:  "",
	}

	got := summits[3]
	if want != got {
		t.Fatal("want: ", spew.Sdump(want), "got:", spew.Sdump(got))
	}
	t.Log(spew.Sdump(got))
}
