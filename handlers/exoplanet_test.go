package handlers

import (
	"bytes"
	"encoding/json"
	"exoplanet-service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddExoplanet(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//post exoplanets API call for add data
	router := gin.Default()
	router.POST("/exoplanets", AddExoplanet)

	exoplanet := models.Exoplanet{
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      5.0,
		Mass:        5.0,
		Type:        models.Terrestrial,
	}

	jsonValue, _ := json.Marshal(exoplanet)
	req, _ := http.NewRequest("POST", "/exoplanets", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Exoplanet
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Planet", response.Name)
	assert.NotEmpty(t, response.ID)

	//Get all List API call testing
	router.GET("/exoplanets", ListExoplanets)

	req, _ = http.NewRequest("GET", "/exoplanets", nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetExoplanetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/exoplanets/:id", GetExoplanetByID)

	testID := "test1"
	exoplanets[testID] = models.Exoplanet{
		ID:          testID,
		Name:        "Test Planet",
		Description: "A test planet",
		Distance:    50,
		Radius:      5.0,
		Mass:        5.0,
		Type:        models.Terrestrial,
	}

	req, _ := http.NewRequest("GET", "/exoplanets/"+testID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Exoplanet
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, testID, response.ID)
	assert.Equal(t, "Test Planet", response.Name)
}
