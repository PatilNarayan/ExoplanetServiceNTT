## Exoplanet Service

This is a Golang microservice for managing exoplanet data using the Gin framework. It provides functionality for adding, listing, retrieving, updating, and deleting exoplanets. The service also includes validation for exoplanet data, unit tests, and sorting/filtering options.

## Features

* Add a new exoplanet
* List all exoplanets with sorting and filtering options
* Retrieve an exoplanet by ID
* Update an exoplanet by ID
* Delete an exoplanet by ID
* Estimate fuel for a journey to an exoplanet
* Validate exoplanet data

## Running the Service

run below commad

```
go run main.go
```

## API Endpoints

1. Method - POST  ; Endpoint - "/exoplanets"

   `{ "name": "Planet", "description": "A test planet", "distance": 50, "radius": 1.0, "mass": 5.0, "type": "Terrestrial" } `

   pass this json data to add exoplanet data
2. Method - GET ; Endpoint - "/exoplanets"

   to get exoplanet data
3. Method - GET ; Endpoint - "/exoplanets/:id"

   to get exoplanet specific ID data
4. Method - PUT ; Endpoint - "/exoplanets/:id"
   `{ "name": "Updated Planet", "description": "An updated planet", "distance": 60, "radius": 1.2, "mass": 5.5, "type": "Terrestrial" }`

   pass this json data to update data in put method
5. Method - DELETE ; Endpoint - "/exoplanets/:id"

   to delete specific data
6. Method - GET ; Endpoint - "/exoplanets/:id/fuel?crew=int"

   pass in place of int any value to get fuel estimation
