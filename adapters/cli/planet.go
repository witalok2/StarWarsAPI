package cli

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/witalok2/starwarsplanet/application"
)

func Run(service application.PlanetServiceInterface, action string, planetID uuid.UUID, planetName, planetClimate, planetTerrain string, appearances int) (string, error) {
	var result = ""

	switch action {
	case "create":
		planet, err := service.Create(planetName, planetClimate, planetTerrain, appearances)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Planet ID %s with the name %s has been created", planet.GetID().String(), planet.GetName())

	default:
		res, err := service.Get(planetID)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Planet ID %s\n Name %s", res.GetID(), res.GetName())
	}

	return result, nil
}
