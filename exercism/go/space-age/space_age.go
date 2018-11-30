package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	earthYears := map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return secToEY(seconds) * 1 / earthYears[planet]
}

func secToEY(secs float64) float64 {
	return secs / 31557600
}
