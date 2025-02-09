// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package cal

import (
	"testing"
	"time"
)

func TestIsApplicable(t *testing.T) {
	zone1 := time.FixedZone("test1", -5)
	zone2 := time.FixedZone("test2", 3)

	tests := []struct {
		c    Calendar
		loc  *time.Location
		want bool
	}{
		{Calendar{}, time.UTC, true},
		{Calendar{Locations: []*time.Location{time.UTC}}, time.UTC, true},
		{Calendar{Locations: []*time.Location{zone1}}, time.UTC, false},
		{Calendar{Locations: []*time.Location{zone1, zone2}}, zone2, true},
	}

	for _, test := range tests {
		got := test.c.IsApplicable(test.loc)
		if got != test.want {
			t.Errorf("got: %t; want: %t (%v)", got, test.want, test.c)
		}
	}
}

func TestAddHoliday(t *testing.T) {
	c := Calendar{}
	c.AddHoliday(&Holiday{})
	c.AddHoliday(&Holiday{})
	if len(c.Holidays) != 2 {
		t.Errorf("got: %d; want: 2", len(c.Holidays))
	}
}

func TestIsHoliday(t *testing.T) {
	zone1 := time.FixedZone("test1", -5)
	zone2 := time.FixedZone("test2", 3)

	hol := &Holiday{
		Type:  ObservancePublic,
		Month: time.July,
		Day:   4,
		Observed: []AltDay{
			{Day: time.Saturday, Offset: -1},
			{Day: time.Sunday, Offset: 1},
		},
		Func: CalcDayOfMonth,
	}

	cachedCalendar := &Calendar{
		Holidays:  []*Holiday{hol},
		Cacheable: true,
	}
	CacheMaxSize = 2
	CacheEvictSize = 1

	tests := []struct {
		c       *Calendar
		date    time.Time
		wantAct bool
		wantObs bool
		wantHol *Holiday
	}{
		{&Calendar{}, time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC), false, false, nil},
		{&Calendar{Holidays: []*Holiday{hol},
			Locations: []*time.Location{zone1, zone2}},
			time.Date(2015, 7, 4, 12, 30, 0, 0, time.UTC), false, false, nil},
		{&Calendar{Holidays: []*Holiday{hol},
			Locations: []*time.Location{zone1, zone2}},
			time.Date(2015, 7, 4, 12, 30, 0, 0, zone1), true, false, hol},
		{&Calendar{Holidays: []*Holiday{hol},
			Locations: []*time.Location{zone1, zone2}},
			time.Date(2015, 7, 3, 12, 30, 0, 0, zone2), false, true, hol},
		{&Calendar{Holidays: []*Holiday{hol},
			Locations: []*time.Location{zone1, zone2}},
			time.Date(2015, 8, 4, 12, 30, 0, 0, zone2), false, false, nil},

		{cachedCalendar, time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC), false, false, nil},
		{cachedCalendar, time.Date(2000, 1, 1, 12, 30, 0, 0, time.UTC), false, false, nil},
		{cachedCalendar, time.Date(2015, 7, 4, 12, 30, 0, 0, time.UTC), true, false, hol},
		{cachedCalendar, time.Date(2015, 7, 4, 12, 30, 0, 0, time.UTC), true, false, hol},
		{cachedCalendar, time.Date(2015, 7, 3, 12, 30, 0, 0, time.UTC), false, true, hol},
		{cachedCalendar, time.Date(2015, 7, 3, 12, 30, 0, 0, time.UTC), false, true, hol},
	}

	for i, test := range tests {
		gotAct, gotObs, gotHol := test.c.IsHoliday(test.date)
		if gotAct != test.wantAct || gotObs != test.wantObs || gotHol != test.wantHol {
			t.Errorf("gotAct: %t, wantAct: %t; gotObs: %t, wantObs: %t; gotHol: %v, wantHol: %v (%d)",
				gotAct, test.wantAct, gotObs, test.wantObs, gotHol, test.wantHol, i)
		}
	}
}
