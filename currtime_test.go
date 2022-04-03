package currtime

import (
	"fmt"
	"testing"
	"time"
)

func testTimeZone(t *testing.T, timezone string) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		t.Fatal(err)
	}

	tt := time.Now().In(loc)
	pt, err := CurTime(timezone)
	if err != nil {
		t.Fatal(err)
	}

	cur_diff := tt.Sub(pt)
	min_diff, _ := time.ParseDuration("1ms")
	if cur_diff > min_diff {
		t.Fatal("package provide incorrect time")
	}

	ttLocation := fmt.Sprint(tt.Location())
	ptLocation := fmt.Sprint(pt.Location())

	if ttLocation != ptLocation {
		t.Fatal("package provide incorrent timezone/location")
	}
}

func TestDefaultTimeZone(t *testing.T) {
	testTimeZone(t, DefaultLocation)
}

func TestLosAngeles(t *testing.T) {
	testTimeZone(t, "America/Los_Angeles")
}
