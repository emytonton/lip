package expenses

import "errors"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var result []Record
	for _, record := range in {
		if predicate(record) {
			result = append(result, record)
		}
	}
	return result
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filtered := Filter(in, ByDaysPeriod(p))
	var total float64
	for _, record := range filtered {
		total += record.Amount
	}
	return total
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	anyInCategory := false
	for _, record := range in {
		if record.Category == c {
			anyInCategory = true
			break
		}
	}

	if !anyInCategory {
		return 0, errors.New("unknown category")
	}

	filtered := Filter(in, func(r Record) bool {
		return r.Category == c && r.Day >= p.From && r.Day <= p.To
	})

	var total float64
	for _, record := range filtered {
		total += record.Amount
	}

	return total, nil
}