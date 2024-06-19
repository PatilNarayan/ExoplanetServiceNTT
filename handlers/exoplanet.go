package handlers

import (
	"exoplanet-service/models"
	"exoplanet-service/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var exoplanets = make(map[string]models.Exoplanet)

func AddExoplanet(c *gin.Context) {
	var planet models.Exoplanet
	if err := c.ShouldBindJSON(&planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"AddExoplanetError": err.Error()})
		return
	}

	planet.ID = xid.New().String()

	if err := utils.ValidateExoplanet(planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"AddExoplanetError": err.Error()})
		return
	}

	exoplanets[planet.ID] = planet
	c.JSON(http.StatusOK, planet)
}

func ListExoplanets(c *gin.Context) {
	var planets []models.Exoplanet
	for _, planet := range exoplanets {
		planets = append(planets, planet)
	}
	c.JSON(http.StatusOK, planets)
}

func GetExoplanetByID(c *gin.Context) {
	id := c.Param("id")
	planet, ok := exoplanets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"GetExoplanetByIDError": "Exoplanet not found"})
		return
	}
	c.JSON(http.StatusOK, planet)
}

func UpdateExoplanet(c *gin.Context) {
	id := c.Param("id")
	_, ok := exoplanets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"UpdateExoplanetError": "Exoplanet not found"})
		return
	}

	var planet models.Exoplanet
	if err := c.ShouldBindJSON(&planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"UpdateExoplanetError": err.Error()})
		return
	}
	planet.ID = id

	if err := utils.ValidateExoplanet(planet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"UpdateExoplanetError": err.Error()})
		return
	}

	exoplanets[id] = planet
	c.JSON(http.StatusOK, planet)
}

func DeleteExoplanet(c *gin.Context) {
	id := c.Param("id")
	_, ok := exoplanets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"DeleteExoplanetError": "Exoplanet not found"})
		return
	}
	delete(exoplanets, id)
	c.Status(http.StatusNoContent)
}

func FuelEstimation(c *gin.Context) {
	id := c.Param("id")
	planet, ok := exoplanets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"FuelEstimationError": "Exoplanet not found"})
		return
	}

	crewCapacity, err := strconv.Atoi(c.Query("crew"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"FuelEstimationError": "Invalid crew capacity"})
		return
	}

	var gravity float64
	if planet.Type == models.GasGiant {
		gravity = 0.5 / (planet.Radius * planet.Radius)
	} else {
		gravity = planet.Mass / (planet.Radius * planet.Radius)
	}

	fuel := float64(planet.Distance) / (gravity * gravity) * float64(crewCapacity)
	c.JSON(http.StatusOK, gin.H{"fuel": fuel})
}
