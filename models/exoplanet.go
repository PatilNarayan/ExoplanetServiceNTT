package models

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    int           `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}
