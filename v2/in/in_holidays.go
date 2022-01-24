// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).

// Package ie provides holiday definitions for the Republic Of Ireland.
package in

import (
	"time"

	"github.com/devechelon/cal/v2"
	"github.com/devechelon/cal/v2/aa"
)

var (
	// NewYear represents New Year's Day on 1-Jan
	NewYear = aa.NewYear.Clone(&cal.Holiday{})

	// GoodFriday represents Good Friday - two days before Easter
	GoodFriday = aa.GoodFriday.Clone(&cal.Holiday{})

	// ChristmasDay represents Christmas Day on 25-Dec
	ChristmasDay = aa.ChristmasDay.Clone(&cal.Holiday{})

	// Pongal represents Pongal 1-13 - 1-17
	Pongal = &cal.Holiday{
		Name:   "Pongal",
		Month:  time.January,
		Day:    17,
		Offset: -3,
		Func:   cal.CalcDayOfMonth,
	}

	// RepublicDay represents RepublicDay 1-26
	RepublicDay = &cal.Holiday{
		Name:   "Republic Day",
		Month:  time.January,
		Day:    26,
		Func:   cal.CalcDayOfMonth,
	}

	//AmbedkarJayanti represents AmbedkarJayanti 04-14
	AmbedkarJayanti = &cal.Holiday{
		Name:   "Ambedkar Jayanti",
		Month:  time.April,
		Day:    14,
		Func:   cal.CalcDayOfMonth,
	}

	// TamilNewYear represents TamilNewYear 1-26
	TamilNewYear = &cal.Holiday{
		Name:   "Tamil New Year",
		Month:  time.April,
		Day:    15,
		Offset: -1,
		Func:   cal.CalcDayOfMonth,
	}

	// Ramazan
	Ramazan = &cal.Holiday{
		Name:   "Ramazan / Idu'l Fitr",
		Month:  time.May,
		Day:    3,
		Offset: -1,
		Func:   cal.CalcDayOfMonth,
	}

	// Muharram
	Muharram = &cal.Holiday{
		Name:   "Muharram",
		Month:  time.August,
		Day:    9,
		Func:   cal.CalcDayOfMonth,
	}

	// IndependenceDay represents IndependenceDay 1-26
	IndependenceDay = &cal.Holiday{
		Name:   "Independence Day",
		Month:  time.August,
		Day:    15,
		Func:   cal.CalcDayOfMonth,
	}

	// Krishna Jayenthi
	KrishnaJayenthi = &cal.Holiday{
		Name:   "Krishna Jayenthi",
		Month:  time.August,
		Day:    19,
		Func:   cal.CalcDayOfMonth,
	}

	// Vinayakar Chathurthi
	VinayakarChathurthi = &cal.Holiday{
		Name:   "Vinayakar Chathurthi",
		Month:  time.August,
		Day:    31,
		Func:   cal.CalcDayOfMonth,
	}

	//Ayutha Pooja represents GandhiJayanti 10-04
	AyuthaPooja = &cal.Holiday{
		Name:   "Ayutha Pooja",
		Month:  time.October,
		Day:    04,
		Func:   cal.CalcDayOfMonth,
	}

	//Vijaya Dasami
	VijayaDasami = &cal.Holiday{
		Name:   "Vijaya Dasami",
		Month:  time.October,
		Day:    05,
		Func:   cal.CalcDayOfMonth,
	}

	//Deepavali
	Deepavali = &cal.Holiday{
		Name:   "Deepavali",
		Month:  time.October,
		Day:    24,
		Func:   cal.CalcDayOfMonth,
	}


	// RamaNavami represents RamaNavami 11-04
	//RamaNavami = &cal.Holiday{
	//	Name:   "Rama Navami",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}

	// Eid al-Fitr represents Eid al-Fitr 11-04
	//EidalFitr = &cal.Holiday{
	//	Name:   "Eid al-Fitr",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}

	// Eid al-Adha represents Eid al-Fitr 11-04
	//EidalAdha = &cal.Holiday{
	//	Name:   "Eid al-Adha",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}

	// Dussehra represents Dussehra 11-04
	//Dussehra = &cal.Holiday{
	//	Name:   "Dussehra",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}

	// Mawlid represents Mawlid 11-04
	//Mawlid = &cal.Holiday{
	//	Name:   "Mawlid",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}

	// Diwali represents Diwali 11-04
	//Diwali = &cal.Holiday{
	//	Name:   "Diwali",
	//	Month:  time.April,
	//	Day:    15,
	//	Offset: -1,
	//	Func:   cal.CalcDayOfMonth,
	//}


	// Holidays provides a list of the standard national holidays
	Holidays = []*cal.Holiday{
		NewYear,
		ChristmasDay,
		GoodFriday,
		Pongal,
		RepublicDay,
		AmbedkarJayanti,
		TamilNewYear,
		Ramazan,
		Muharram,
		IndependenceDay,
		KrishnaJayenthi,
		VinayakarChathurthi,
		AyuthaPooja,
		VijayaDasami,
		Deepavali,
	}
)
