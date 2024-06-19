package routes

import (
	"exoplanet-service/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/exoplanets", handlers.AddExoplanet)
	router.GET("/exoplanets", handlers.ListExoplanets)
	router.GET("/exoplanets/:id", handlers.GetExoplanetByID)
	router.PUT("/exoplanets/:id", handlers.UpdateExoplanet)
	router.DELETE("/exoplanets/:id", handlers.DeleteExoplanet)
	router.GET("/exoplanets/:id/fuel", handlers.FuelEstimation)

	return router
}
