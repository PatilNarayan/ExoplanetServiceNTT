package main

import (
	"exoplanet-service/routes"
)

func main() {
	router := routes.RegisterRoutes()
	router.Run(":8080")
}
