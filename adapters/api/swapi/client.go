package swapi

import (
	"errors"

	"github.com/parnurzeal/gorequest"
)

type Client struct {
	Host       string
	superAgent *gorequest.SuperAgent
}

type ResultPlanet struct {
	Count    int32
	Next     string
	Previous string
	Results  []Planet
}

type Planet struct {
	Name           string
	RotationPeriod string
	OrbitalPeriod  string
	Diameter       string
	Climate        string
	Gravity        string
	Terrain        string
	SurfaceWater   string
	Population     string

	Residents []string
	Films     []string

	Created string
	Edited  string
	URL     string
}

var (
	errInvalid = errors.New("Erro no cliente")
)

func (c *Client) request() (superAgent *gorequest.SuperAgent) {
	if c.superAgent != nil {
		return c.superAgent
	}

	c.superAgent = gorequest.New()
	return c.superAgent
}

func (c Client) get(targetURL string) (superAgent *gorequest.SuperAgent) {
	return c.request().Get(c.Host + targetURL)
}

func (c Client) GetPlanetInSwapiByName(planetName string) (planets ResultPlanet, err error) {
	tempStruct := struct {
		Count    int32  `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name           string `json:"name"`
			RotationPeriod string `json:"rotation_period"`
			OrbitalPeriod  string `json:"orbital_period"`
			Diameter       string `json:"diameter"`
			Climate        string `json:"climate"`
			Gravity        string `json:"gravity"`
			Terrain        string `json:"terrain"`
			SurfaceWater   string `json:"surface_water"`
			Population     string `json:"population"`

			Residents []string `json:"residents"`
			Films     []string `json:"films"`

			Created string `json:"created"`
			Edited  string `json:"edited"`
			URL     string `json:"url"`
		} `json:"results"`
	}{}

	response, _, errs := c.get("planets/?search=" + planetName).EndStruct(&tempStruct)
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	if response.StatusCode != 200 {
		err = errInvalid
		return
	}

	planets.Count = tempStruct.Count
	planets.Next = tempStruct.Next
	planets.Previous = tempStruct.Previous

	for _, result := range tempStruct.Results {
		planets.Results = append(planets.Results, Planet{
			Name:           result.Name,
			RotationPeriod: result.RotationPeriod,
			OrbitalPeriod:  result.OrbitalPeriod,
			Diameter:       result.Diameter,
			Climate:        result.Climate,
			Gravity:        result.Gravity,
			Terrain:        result.Terrain,
			SurfaceWater:   result.SurfaceWater,
			Population:     result.Population,

			Residents: result.Residents,
			Films:     result.Films,

			Created: result.Created,
			Edited:  result.Edited,
			URL:     result.URL,
		})
	}

	return
}

func (c Client) ListPlanetInSwapi() (planets ResultPlanet, err error) {
	tempStruct := struct {
		Count    int32  `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name           string `json:"name"`
			RotationPeriod string `json:"rotation_period"`
			OrbitalPeriod  string `json:"orbital_period"`
			Diameter       string `json:"diameter"`
			Climate        string `json:"climate"`
			Gravity        string `json:"gravity"`
			Terrain        string `json:"terrain"`
			SurfaceWater   string `json:"surface_water"`
			Population     string `json:"population"`

			Residents []string `json:"residents"`
			Films     []string `json:"films"`

			Created string `json:"created"`
			Edited  string `json:"edited"`
			URL     string `json:"url"`
		} `json:"results"`
	}{}

	response, _, errs := c.get("planets").EndStruct(&tempStruct)
	if len(errs) > 0 {
		err = errs[0]
		return
	}

	if response.StatusCode != 200 {
		err = errInvalid
		return
	}

	planets.Count = tempStruct.Count
	planets.Next = tempStruct.Next
	planets.Previous = tempStruct.Previous

	for _, result := range tempStruct.Results {
		planets.Results = append(planets.Results, Planet{
			Name:           result.Name,
			RotationPeriod: result.RotationPeriod,
			OrbitalPeriod:  result.OrbitalPeriod,
			Diameter:       result.Diameter,
			Climate:        result.Climate,
			Gravity:        result.Gravity,
			Terrain:        result.Terrain,
			SurfaceWater:   result.SurfaceWater,
			Population:     result.Population,

			Residents: result.Residents,
			Films:     result.Films,

			Created: result.Created,
			Edited:  result.Edited,
			URL:     result.URL,
		})
	}

	return
}
