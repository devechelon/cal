// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

package in

import (
	"testing"
	"time"

	"github.com/devechelon/cal/v2"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	tests := []struct {
		h    *cal.Holiday
		y    int
		date time.Time
	}{
		{NewYear, 2020, d(2020, 1, 1)},
		{NewYear, 2021, d(2021, 1, 1)},
		{NewYear, 2022, d(2022, 1, 1)},

		{ChristmasDay, 2020, d(2020, 12, 25)},
		{ChristmasDay, 2021, d(2021, 12, 25)},
		{ChristmasDay, 2022, d(2022, 12, 25)},

	}

	for _, test := range tests {
		gotAct, _ := test.h.Calc(test.y)
		if !gotAct.Equal(test.date) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.date.String())
		}
	}
}
