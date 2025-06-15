package space

import "math"

type Planet string


func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / 31557600.0

	
	orbitalPeriods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	period, exists := orbitalPeriods[planet]
	if !exists {
		return -1.0
	}

	age := earthYears / period
	return math.Round(age*100) / 100
}