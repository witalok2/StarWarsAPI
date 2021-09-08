package dto

import (
	"github.com/google/uuid"
	"github.com/witalok2/starwarsplanet/application"
)

type Planet struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Climate string    `json:"climate"`
	Terrain string    `json:"terrain"`
}

func NewPlanet() *Planet {
	return &Planet{}
}

func (p *Planet) Bind(planet *application.Planet) (*application.Planet, error) {
	if p.ID.String() != "" {
		planet.ID = p.ID
	}

	planet.Name = p.Name
	planet.Climate = p.Climate
	planet.Terrain = p.Terrain

	_, err := planet.IsValid()
	if err != nil {
		return &application.Planet{}, err
	}

	return planet, nil
}
