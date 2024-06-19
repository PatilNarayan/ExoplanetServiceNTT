package utils

import (
	"errors"
	"exoplanet-service/models"
)

func ValidateExoplanet(planet models.Exoplanet) error {
	if planet.Name == "" || planet.Description == "" || planet.Distance <= 10 || planet.Distance >= 1000 || planet.Radius <= 0.1 || planet.Radius >= 10 {
		return errors.New("invalid exoplanet data")
	}
	if planet.Type == models.Terrestrial && (planet.Mass <= 0.1 || planet.Mass >= 10) {
		return errors.New("invalid mass for terrestrial exoplanet")
	}
	return nil
}
