package booking

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layouts := []string{
		"1/2/2006 3:04:05 PM",
		"1/2/2006 15:04:05",
		"January 2, 2006 3:04:05 PM",
		"January 2, 2006 15:04:05",
		"January 2, 2006 3:04:05",
		"Monday, January 2, 2006 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}

	panic(fmt.Sprintf("Could not parse date string: %s", date))
}

func HasPassed(date string) bool {
	t := Schedule(date)
	return time.Now().After(t)
}

func IsAfternoonAppointment(date string) bool {
	t := Schedule(date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s.",
		t.Weekday().String(),
		t.Format("January 2, 2006, at 15:04"),
	)
}

func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
