package weather

var CurrentCondition string

var CurrentLocation string

func Forecast(city, condition string) string {
	CurrentLocation = city
	CurrentCondition = condition
	return city + " - current weather condition: " + condition
}
