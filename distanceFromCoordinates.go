package distance_util

import (
	"errors"
	"math"
)

// Coordinates struct to map latitude and longitude
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// GetDistanceFromCoordinates function will accept two parameters
// Origin Coordinates and Destination Coordinates
// and will convert the latitude and longitude coordinates to distance using mathematical calculations
func GetDistanceFromCoordinates(origin, destination Coordinates) (float64, error) {

	if origin.Latitude < 0 || origin.Longitude < 0 {
		return 0, errors.New("origin latitude/longitude must be positive")
	}

	if destination.Latitude < 0 || destination.Longitude < 0 {
		return 0, errors.New("destination latitude/longitude must be positive")
	}

	if origin.Latitude > 90 || origin.Longitude > 180 {
		return 0, errors.New("origin latitude must be less than 90 and origin longitude must be less than 180")
	}

	if destination.Latitude > 90 || destination.Longitude > 180 {
		return 0, errors.New("destination latitude must be less than 90 and destination longitude must be less than 180")
	}

	originLatitudeInRadians := math.Pi * origin.Latitude / 180
	destinationLatitudeInRadians := math.Pi * destination.Latitude / 180

	theta := origin.Longitude - destination.Longitude
	radiansTheta := math.Pi * theta / 180

	dist := math.Sin(originLatitudeInRadians)*math.Sin(destinationLatitudeInRadians) + math.Cos(originLatitudeInRadians)*math.Cos(destinationLatitudeInRadians)*math.Cos(radiansTheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	dist = dist * 1.609344

	return dist, nil
}
