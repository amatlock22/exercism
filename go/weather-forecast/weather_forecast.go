// Package weather contains the forecast for a city.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the city for the weather.
var CurrentLocation string

// Forecast takes in the city and weather condition to return a string describing the weather for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
