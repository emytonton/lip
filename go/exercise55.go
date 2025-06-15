package clock

import "fmt"

const (
	minutesInHour = 60
	minutesInDay  = 24 * minutesInHour
)


type Clock int


func New(h, m int) Clock {
	
	totalMinutes := h*minutesInHour + m

	totalMinutes = (totalMinutes%minutesInDay + minutesInDay) % minutesInDay

	return Clock(totalMinutes)
}


func (c Clock) String() string {
	hours := c / minutesInHour
	minutes := c % minutesInHour
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}


func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}


func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}
