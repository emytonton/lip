package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}